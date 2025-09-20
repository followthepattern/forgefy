import 'package:adapticc/app/account/lib/query_properties.dart';
import 'package:adapticc/graphql/list_response.dart';
import 'package:flutter/material.dart';
import 'package:graphql/client.dart';

abstract class ListQuery<T> extends ChangeNotifier {
  ListQuery({
    required this.graphQLClient,
    required this.queryProperties,
  });

  final GraphQLClient graphQLClient;
  final QueryProperties queryProperties;

  bool _isLoading = false;
  bool _hasError = false;
  ListResponse<T> _listResponse = ListResponse.empty();

  bool get isLoading => _isLoading;
  bool get hasError => _hasError;
  int get count => _listResponse.count;
  List<T> get data => _listResponse.data;

  void init() {
    queryProperties.addListener(_onQueryPropertiesChanged);
    _fetch();
  }

  @override
  void dispose() {
    queryProperties.removeListener(_onQueryPropertiesChanged);
    super.dispose();
  }

  ListResponse<T> parseResult(Map<String, dynamic> data);
  String get query;

  Future<void> _fetch() async {
    _hasError = false;
    _isLoading = true;
    notifyListeners();

    try {
      final result = await graphQLClient.query(
        QueryOptions(
          document: gql(query),
          variables: {
            'pageSize': queryProperties.pageSize,
            'page': queryProperties.currentPage,
            'search': queryProperties.filter ?? '',
            'orderBy': queryProperties.orderBy?.toJson(),
          },
        ),
      );
      if (result.data == null) {
        _isLoading = false;
        notifyListeners();
        return;
      }
      _listResponse = parseResult(result.data!);
      _isLoading = false;
      notifyListeners();
    } catch (e) {
      _hasError = true;
      _isLoading = false;
      notifyListeners();
      return;
    }
  }

  void _onQueryPropertiesChanged() {
    _fetch();
    notifyListeners();
  }
}
