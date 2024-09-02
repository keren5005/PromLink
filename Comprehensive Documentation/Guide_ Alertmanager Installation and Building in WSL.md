# Guide: Alertmanager Installation and Building in WSL

### Authors:
- Sharona Seleri
- Yehonatan Ailon
- Keren Cohen
- Amit Sindani

### Requirements:
- **Operating System**: Windows 10 (V2004 or higher) / 11
  - *Note*: You don’t need to upgrade your operating system to Windows 11 from Windows 10 to install Prometheus server in WSL.
- **WSL Version**: Version 2 installed in Windows
  - *Note*: If you don’t have WSL installed in your Windows or need to upgrade your current WSL version from V1 to V2 – follow [this guide](https://learn.microsoft.com/en-us/windows/wsl/install).

### Installation Steps:

1. **Installing the latest versions of GO, Node.js, and npm.**
   - *Notes*:
     - Perform the installations through the WSL terminal in the path of the home folder (e.g., `/home/amitz`).
     - If certain versions of GO, Node.js or npm already exist in your WSL, our recommendation is to delete them first (this can avoid duplicate download issues).
     - It is also recommended to delete NVM in WSL as well if it exists.

2. **Installing GO:**
   - Download GO from [here](https://go.dev/doc/install).
   - Follow the installation procedure under "Linux".
   - Execute the following command every time you enter the terminal (otherwise, GO might not be found when you try to use it):
     ```bash
     export PATH=$PATH:/usr/local/go/bin
     ```

3. **Installing Node.js and npm:**
   - Use the following commands one after the other:
     ```bash
     sudo apt update
     sudo apt upgrade
     sudo apt install nodejs npm
     ```
   - Check their versions after installation.
   - Install NVM in WSL to manage and switch Node.js versions using the commands from [this link](https://linuxbeast.com/blog/how-to-switch-node-js-version-in-wsl-ubuntu/).
  
3. **Clone Alertmanager repository:**
   - Run the following commands:
     ```bash
     git clone https://github.com/prometheus/alertmanager.git
     cd alertmanager
     ```

4. **Build AlertManager:**
   - Run the following command in the `alertmanager` directory:
     ```bash
     make build
     ```

5. **Run AlertManager with the configuration file:**
   - You can download the configuration file from [this GitHub page](https://github.com/prometheus/alertmanager/blob/main/doc/examples/simple.yml).
   - Run AlertManager with the following command:
     ```bash
     ./alertmanager --config.file=<your_file>
     ```

6. **Access the UI:**
   - Enter the UI in your browser at the address: `http://localhost:9093/#/alerts`.
