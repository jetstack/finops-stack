#!/usr/bin/env python
import os
from collections import defaultdict
from urllib.parse import urlparse

from diagrams import Cluster, Diagram, Edge

from diagrams.onprem.monitoring import Prometheus, Grafana
from diagrams.onprem.client import Users
from diagrams.custom import Custom

from diagrams.k8s.compute import Deploy, StatefulSet, Cronjob
from diagrams.k8s.others import CRD
from diagrams.generic.place import Datacenter

from diagrams.k8s.ecosystem import Helm

import yaml

with Diagram("Architecture",
             filename="../content/assets/architecture",
             outformat="png",
             show=False,
             direction="TB"):

    users = Users("Users")

    with Cluster("Cluster"):
        with Cluster("[ns] applications"):
            a1 = Deploy("App1")
            a2 = StatefulSet("App2")
            a3 = Cronjob("App3")
            _ = a1 - Edge(color="transparent") - a2 - Edge(color="transparent")
        with Cluster("[ns] finops-stack"):
            prom = Prometheus("Prometheus")
            graf = Grafana("Grafana")
            kyverno = Custom("Kyverno", os.getcwd() +
                                "/resources/kyverno.png")
            finops_policies = CRD("Finops Policies")
            opencost = Custom("Opencost", os.getcwd() +
                                "/resources/opencost.png")
            finops_dashboards = CRD("Finops Dashboards")
            vpa = CRD("VPA")

            # _ = vpa - Edge(color="transparent") -

        # _ = a3 - Edge(color="transparent") - vpa

        for i in range(0, 0):
            with Cluster(f"[ns] Application {i}"):
                Deploy(f"Application {i}")

        aws = Custom("AWS", os.getcwd()+'/resources/aws.png')
        azure = Custom("Azure", os.getcwd()+'/resources/azure.png')
        gcp = Custom("GCP", os.getcwd()+'/resources/google-cloud.png')
        onprem = Datacenter("On-Prem")

        _ = users >> graf
        _ = graf >> prom
        _ = graf << Edge(style="dashed") << finops_dashboards
        _ = kyverno << Edge(style="dashed") << finops_policies
        _ = prom >> opencost

        for i in [aws, azure, gcp, onprem]:
            _ = opencost << Edge(label="Cost Data", style='dotted') << i


with Diagram("Chart Structure",
             filename="../content/assets/structure",
             outformat="png",
             show=False,
             direction="TB"):

    with open("../chart/Chart.yaml", 'r', encoding='UTF-8') as file:
        chart_yaml = yaml.safe_load(file)

    with Cluster('finops-stack'):
        parent_chart = Helm("finops-stack")

        grouped = defaultdict(list)
        for dependency in chart_yaml['dependencies']:
            parsed_url = urlparse(dependency['repository'])
            grouped[parsed_url.netloc].append(dependency)

        for repo, dept in grouped.items():
            with Cluster(repo):
                for dep in dept:
                    parent_chart >> Helm(dep['name'])
