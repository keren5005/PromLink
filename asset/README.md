# PromLink 

**Team members:** Sharona Seleri, Yehontan Ailon, Keren Cohen, Amit Sindani

## Intro

Prometheus is an open-source monitoring solution that collects and stores metrics data in a time series database.
Alertmanager, a key component of Prometheus, manages alerts by deduplicating, grouping, and routing them to various integrations like email and Slack. However, there are some chat applications, such as Mattermost, that are not supported, leaving organizations using it for internal communication without direct alerts from Prometheus.

Our project aims to enhance Prometheus Alertmanager by integrating it with the chat applications that are still not supported by Prometheus.Our goal was to bridge this gap, allowing Prometheus Alertmanager to send alerts directly to Mattermost channels, thus improving operational efficiency and responsiveness.

![image](https://github.com/keren5005/PromLink/assets/120311888/c8ebd53b-b21b-4a30-914a-e5f564b58daf)

## The Project Document
[The Project Document](https://docs.google.com/document/d/1Z8aP73-qfQTJvO1sbzkbGy7vpMwKQOuJ/edit)

[Demonstration Video](https://drive.google.com/file/d/1Fke3X5ISK8K_aoDAEpY8S_DVMMm4ngh1/view?usp=drive_link)

## Installation Manuals
<br/>[Prometheus Server Installation and Building in WSL](https://docs.google.com/document/d/1b3tM0dMIS8sU3uE8-a8AZPi6YwI0U_NX/edit)

[AlertManager Installation and Building in WSL](https://docs.google.com/document/d/1yZOTlduuzcRc3hGHFiJ0rQJclj3DQkp_/edit).

## Alertmanager

Alertmanager processes alerts from client applications like the Prometheus server. It handles alert deduplication, grouping, and routing to appropriate receiver integrations, such as email, PagerDuty, or OpsGenie. It also supports alert silencing and inhibition. Alertmanager can be installed using precompiled binaries, Docker images, or by compiling the source code manually. Configuration is done via a YAML file, allowing for detailed alert routing, grouping, and inhibition rules.

For high availability, Alertmanager instances can be clustered, ensuring redundancy and reliability. The `amtool` CLI facilitates interaction with the Alertmanager API, allowing for alert management, silencing, and configuration testing. The Alertmanager API, generated via OpenAPI and Go Swagger, provides comprehensive endpoints for alert and configuration management.

## Architecture Diagram

![Alertmanager Architecture](https://raw.githubusercontent.com/keren5005/alertmanager/b0f54f0bc76114b03d81627c714a7d12e4138795/doc/arch.svg)

## Contributed Alertmanagers

### Alertmanager for Metarmost

<div style="background-color:#f0f0f0; padding:10px; border-radius:5px;">
<h4>Introduction</h4>
<p>Here are the Alertmanager setup for Metarmost. Below are the steps to clone the repository, build the project, and run the Alertmanager, as well as instructions for sending alerts using Postman.</p>

<h4>Setup Instructions</h4>
<ol>
<li><strong>Clone the Repository:</strong> First, you need to clone the repository from GitHub. Use the following command:</li>
<pre><code>git clone https://github.com/keren5005/alertmanager.git</code></pre>

<li><strong>Set Environment Variable:</strong> Set the <code>GO15VENDOREXPERIMENT</code> environment variable to 1. This is necessary for building the project.</li>
<pre><code>set GO15VENDOREXPERIMENT=1</code></pre>

<li><strong>Navigate to Alertmanager Directory:</strong> Change your directory to the <code>alertmanager</code> directory.</li>
<pre><code>cd alertmanager</code></pre>

<li><strong>Get Dependencies:</strong> Fetch the required dependencies for the Alertmanager.</li>
<pre><code>go get github.com/prometheus/alertmanager/cmd/...</code></pre>

<li><strong>Build the Alertmanager:</strong> Navigate to the <code>cmd/alertmanager</code> directory and build the project.</li>
<pre><code>cd cmd/alertmanager
go build</code></pre>

<li><strong>Copy Configuration File:</strong> Copy the sample configuration file to use it as your configuration.</li>
<pre><code>cd ../..
copy cmd/alertmanager/alertmanager.yml.sample cmd/alertmanager/alertmanager.yml</code></pre>

<li><strong>Run Alertmanager:</strong> Run the Alertmanager executable.</li>
<pre><code>cmd/alertmanager/alertmanager.exe</code></pre>

<li><strong>Access Alertmanager Web Interface:</strong> Open your web browser and go to: <a href="http://localhost:9093">http://localhost:9093</a></li>
</ol>

<h4>Development Instructions</h4>
<ol>
<li><strong>Open a New Terminal and Navigate to Alertmanager Directory:</strong> Open a new terminal window and navigate to the <code>alertmanager/cmd/alertmanager</code> directory.</li>
<pre><code>cd alertmanager/cmd/alertmanager</code></pre>

<li><strong>Build the Project:</strong> Build the Alertmanager project.</li>
<pre><code>go build</code></pre>
</ol>

<h4>Sending Alerts with Postman</h4>
<ol>
<li><strong>POST Request to Trigger an Alert:</strong> Send a POST request to <a href="[http://localhost:9093/api/v2/alerts](http://localhost:9093/api/v2/alerts)">http://localhost:9093/api/v2/alerts</a> with the following JSON body:</li>
<pre><code>[
{
    "status": "firing",
	"labels": {
		"alertname": "$name4",
		"service": "my-service4",
		"severity": "warning4",
		"instance": "$name.example.net4"
	},
	"annotations": {
		"summary": "High latency is high!"
	}
}
]
</code></pre>

<li><strong>Closing the Alert:</strong> Run the following command to stop the Alertmanager:</li>
<pre><code>cmd/alertmanager/alertmanager.exe</code></pre>
</ol>
</div>
