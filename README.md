
# PromLink:  Prometheus Alert Integration for Mattermost and RocketChat

## Overview

**PromLink** is a seamless integration tool that connects Prometheus Alertmanager with Mattermost and RocketChat channels, enabling organizations to receive real-time alerts directly in their internal communication platforms. Organizations using Mattermost and RocketChat for internal communication often face challenges in receiving direct alerts from Prometheus Alertmanager. The lack of a straightforward integration mechanism hampers timely responses to critical events. PromLink provides a direct integration between Prometheus Alertmanager and Mattermost/RocketChat, allowing alerts to flow seamlessly into designated channels. This integration ensures that teams can quickly respond to alerts and maintain robust communication during incidents.

## Demo Video

Watch the demo video to see PromLink in action:

[![Watch the video](updated_video_thumbnail.png)](https://www.youtube.com/watch?v=BpbbXHgIpM0)


## What is Prometheus?

Prometheus is an open-source monitoring and alerting toolkit designed for reliability and scalability. Key features include:

- **Time-Series Data**: Efficiently stores and queries metrics over time.
- **PromQL**: A powerful query language for flexible data retrieval.
- **Flexible Collection**: Uses a pull model to collect metrics from configured targets.
- **Built-In Alerting**: Integrates with various notification systems to handle alerts.

## What is Prometheus Alertmanager?

Prometheus Alertmanager is responsible for managing alerts generated by Prometheus. It offers:

- **Alert Management**: Handles alerts from Prometheus and client applications.
- **Flexible Routing**: Routes alerts based on custom labels.
- **Silencing & Inhibition**: Suppresses unnecessary alerts to prevent alert storms.
- **Notification Integrations**: Supports integrations with email, Slack, PagerDuty, and more.
- **Grouping**: Clusters related alerts to reduce notification fatigue.

## What are Mattermost and RocketChat?

Mattermost and RocketChat are open-source messaging platforms tailored for team collaboration. They provide:

- **Real-Time Communication**: Facilitate instant messaging and discussions.
- **Customization**: Highly flexible with plugins and integrations.
- **Deployment Options**: Available as self-hosted or cloud-based solutions.
- **Integration Capabilities**: Seamlessly connect with various tools for streamlined workflows.

## Comprehensive Documentation

The project includes comprehensive installation and configuration guides to assist users in setting up the integration. This documentation is designed to fill gaps found in existing resources and provide clear, step-by-step instructions.

- [Prometheus Server Installation Guide](./Prometheus_Server_Installation_Guide.md)
- [Guide: Alertmanager Installation and Building](./Guide_Alertmanager_Installation_and_Building.md)


## Architecture overview

![image](https://github.com/user-attachments/assets/66b8c63b-b250-4b9a-b4d9-2a74374bc62e)

## Installation

To install PromLink, follow the steps below:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/keren5005/PromLink.git
   ```

2. **Navigate to the Directory**:
   ```bash
   cd PromLink
   ```

3. **Follow the Installation Guide**: Detailed instructions are available in the `docs/INSTALL.md` file.

## Contributing

We welcome contributions from the community! Please see our `CONTRIBUTING.md` file for more information on how to get involved.

## Acknowledgments

We would like to extend our gratitude to Dr. Hadar Binsky for guiding us throughout this project. Special thanks to all contributors and community members who supported us.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
