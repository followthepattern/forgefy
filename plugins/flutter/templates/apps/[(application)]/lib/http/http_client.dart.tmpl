import 'dart:convert';

import 'package:http/http.dart' as http;

class HttpClient {
  Future<Map<String, dynamic>> post({
    required String url,
    required Map<String, dynamic> body,
    Map<String, String> headers = const {},
  }) async {
    final response = await http.post(
      Uri.parse(url),
      body: jsonEncode(body),
      headers: {
        'Content-Type': 'application/json',
        ...headers,
      },
    );

    if (response.statusCode != 200) {
      throw HttpClientException(
        statusCode: response.statusCode,
        message: response.body,
      );
    }

    return jsonDecode(response.body);
  }
}

class HttpClientException implements Exception {
  HttpClientException({required this.statusCode, required this.message});

  final int statusCode;
  final String? message;
}
