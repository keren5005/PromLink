
# Guide: Prometheus Server Installation and Building in WSL

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
     \`\`\`bash
     export PATH=$PATH:/usr/local/go/bin
     \`\`\`

3. **Installing Node.js and npm:**
   - Use the following commands one after the other:
     \`\`\`bash
     sudo apt update
     sudo apt upgrade
     sudo apt install nodejs npm
     \`\`\`
   - Check their versions after installation.
   - Install NVM in WSL to manage and switch Node.js versions using the commands from [this link](https://linuxbeast.com/blog/how-to-switch-node-js-version-in-wsl-ubuntu/).

4. **Download Prometheus Server:**
   - Run the following command:
     \`\`\`bash
     git clone https://github.com/prometheus/prometheus.git
     cd prometheus
     \`\`\`

5. **Building Prometheus:**
   - Run the following command from the Prometheus root folder:
     \`\`\`bash
     make build
     \`\`\`
   - There may be errors during this stage, often due to multiple GO installations in different locations in WSL. Ensure there are no duplicates in GO, Node.js, or npm installations.
   - At the end of the process, ensure that there is an executable file named `prometheus` in the Prometheus folder.

6. **Running Prometheus:**
   - Before running the Prometheus executable, ensure that the file `prometheus.yml` is transferred to the Prometheus root folder.
   - The file can be downloaded from [this GitHub page](https://github.com/prometheus/prometheus/blob/main/documentation/examples/prometheus.yml).
   - Run the server using the following command:
     \`\`\`bash
     ./prometheus --config.file=prometheus.yml
     \`\`\`
   - The UI can be accessed in a browser at the address: `localhost:9090`
