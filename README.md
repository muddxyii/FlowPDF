# Project Setup and Development for FlowPDF
**FlowPDF** is an application built using the **Fyne** framework in **Go**. It is specifically designed to help manage AnyBackflow's specific form PDFs efficiently. The application focuses on simplifying the handling, organization, and filling of these forms, ensuring a streamlined user experience.

## Features of FlowPDF
- **PDF Management**: Easily organize and work with AnyBackflow-specific forms.
- **User-Friendly UI**: Built with **Fyne**, FlowPDF provides a modern, responsive, and platform-independent UI.
- **Go-Based Development**: Ensures a lightweight and high-performance application.
- **Automated Reloading**: Integrated with `air` for hot reloading during development, speeding up the iteration process.

---

## Running Air (Hot Reloading for Go)
To use `air` for hot reloading during development, follow these steps:
1. **Install Air (if not already installed)**: Run the following command to install `air`:

    ```bash
    go install github.com/cosmtrek/air@latest
    ```

2. **Run Air**:
   Use the following command to run `air`:
    ```bash
    $(go env GOPATH)/bin/air
    ```

Air will automatically watch for code changes and restart the application.

---

## Additional Notes

- Make sure your `GOPATH` is correctly set up.
- Ensure `air` is installed in the specified `$(go env GOPATH)/bin` directory.

For more details about `air`, refer to the [official GitHub repository](https://github.com/cosmtrek/air).

---