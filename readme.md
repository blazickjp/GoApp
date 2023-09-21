# 🚀 Python Error Collector for LLMs 🚀

Hello there, adventurous coder! 😎 Are you tired of those pesky Python errors ruining your day? 
Ever wished you could teach your Language Learning Model (LLM) to understand _your_ specific coding mistakes? 
Well, guess what? Your wish has just come true! 🌈

## 🌟 Features

- **Real-Time Error Capturing**: Automatically captures Python errors as they happen. 🐞
- **Codebase Snapshots**: Takes a snapshot of the relevant code files and saves them in a user-specific directory. 📸
- **Shell Flexibility**: Works with both Bash and Zsh. 🐚
- **Systray Integration**: A handy systray menu for quick interactions. 🍱

## 📍 Where Does It Save the Files?

- On macOS: The snapshot and error files get saved to `~/Library/Application Support/PythonCapture`.
- On Windows: The files will be saved in a corresponding AppData folder.

## 🛠 Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/YourUsername/PythonErrorCollectorForLLMs.git
    ```
2. Navigate into the project directory:
    ```bash
    cd PythonErrorCollectorForLLMs
    ```
3. Install the required Go packages (make sure Go is installed):
    ```bash
    go get -u github.com/getlantern/systray
    go get -u github.com/google/uuid
    ```
4. Build the Go project:
    ```bash
    go build
    ```
    
## 🚀 Usage

1. Start the application:
    ```bash
    ./your_compiled_binary
    ```
2. Use the systray menu to:
    - Set your preferred shell 🐚
    - Add or remove an alias 📛

3. Run your Python scripts as you normally would. If an error occurs, the application will capture it along with a snapshot of your codebase.

4. The collected data will be used to fine-tune your LLM. 🎯

## 🤝 Contributing

Feel free to fork, open a pull request, or submit issues. All contributions are welcomed! 🤗

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎉 Final Words

Happy coding! May your Python scripts be ever error-free and your LLMs ever smarter! 🥳
