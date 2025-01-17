# Forgefy

[Forgefy.io](https://forgefy.io/) generates applications effortlessly.

This project aims to reduce the time and cost of software development by generating code using the simplest possible configuration input. The motivation behind Forgefy is to enable building, iterating, and deploying a production-ready application with minimal resource commitment.

## Ideal Use Cases

Forgefy is perfect for:
- **Startups** to build MVPs quickly.
- **Small businesses** to iterate over different approaches.
- **Developers** who want to test their ideas.
- **Anyone** looking to create a cloud-native product.

## Usage  

Forgefy is primarily designed to be used via the [Forgefy.io](https://forgefy.io/) platform. Simply visit the website, define your application following the syntac, and generate it effortlessly.  

The generated code will include a README file specifically tailored to the application, providing clear instructions and details for that particular app.

For those who prefer a local workflow, Forgefy can also be run as a CLI tool:
```bash
go build cmd/forgefy && ./forgefy 
```

## Why Generate Code Instead of Using Frameworks?

While many frameworks aim to lower the cost of building applications by providing general-purpose components, they often result in unmanageable codebases. This happens because frameworks introduce **constraints** designed to handle common behaviors—but there are always exceptions.

Frameworks solve for flexibility during runtime, which often leads to inefficient code, slow queries, and the misuse of uncommon language features like decorators, specific file naming conventions, or flags.

In contrast, code generation does introduce some constraints, but they remain confined to the code generator itself and don’t appear in the application code.

Frameworks also tend to enforce specific architectures, creating a culture where deviating from the prescribed methods feels like "hacking the system."

Generated code, on the other hand, is inherently more scalable because it avoids these runtime constraints.

## AI vs. Template-Generated Code

Generating code with AI is risky. While the capabilities of Gen AI (Generative AI) are intriguing, generating an entire application source code remains a significant challenge.

### Why Not Build Applications Solely With Gen AI?
1. **Inconsistent Outputs**: The same prompt can produce slightly different results, even with the same model, making debugging difficult.
2. **Limited Expression**: It’s hard to convey all necessary details using prompts alone.
3. **Expensive Training**: Modifying behavior through model training is resource-intensive.
4. **High Iteration Costs**: Iterating over variations is more costly compared to template-based code generation.

### Where Forgefy Uses AI

Forgefy employs AI for less critical but time-consuming tasks, such as:
- Generating realistic test data (e.g., book titles, product names, or types).
- Creating themes, colors, and icons.
- Translating the UI into different languages.

By focusing AI on these areas, Forgefy keeps critical tasks predictable and efficient.

## Contribution

The project is not open for contributions yet, but this will change soon. Feedback via issues is more than welcome and greatly appreciated!