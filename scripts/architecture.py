#!/usr/bin/env python
from diagrams import Cluster, Diagram, Edge
import os
from diagrams.onprem.monitoring import Prometheus, Grafana
from diagrams.onprem.client import Users
from diagrams.custom import Custom

from diagrams.k8s.compute import Deploy, StatefulSet, Cronjob
from diagrams.k8s.others import CRD
from diagrams.generic.place import Datacenter

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
