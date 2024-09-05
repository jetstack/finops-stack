# Policy Enforcement

Configuring policy enforcement for FinOps in a cloud-native environment, particularly within Kubernetes, brings several key benefits that help organisations optimise their cloud spending while ensuring compliance and governance. FinOps focuses on creating financial accountability and managing the costs of cloud services, and integrating policy enforcement tools such as Kyverno (used within the finops-stack) into this process helps streamline cost management by enforcing consistent practices across teams and workloads. These policies can govern how resources are used, limiting wasteful spending, enforcing quotas, and ensuring that teams adhere to best practices like right-sizing workloads, using cost-effective storage, or cleaning up unused resources. By applying policies automatically, organisations can avoid uncontrolled cloud expenses and maintain financial discipline at scale.

Incorporating policy enforcement in FinOps also provides transparency and accountability. FinOps teams often work closely with engineering teams to manage cloud usage, but the complexity of cloud-native environments makes it difficult to track which workloads are driving costs. Policies can enforce tagging standards, which make it easier to identify which teams, projects, or environments are responsible for resource consumption. This visibility allows for more accurate cost allocation and reporting, helping teams understand their spending patterns and optimise them. Policy enforcement can also ensure that resources deployed in production adhere to budget constraints, preventing accidental overspending by stopping deployments that exceed predefined cost thresholds or enforcing the use of reserved or spot instances where applicable.

Additionally, policy enforcement helps maintain a balance between agility and governance. One of the core principles of FinOps is to provide engineering teams with autonomy while still enforcing cost control. Policies allow organisations to set guardrails without hindering developer velocity, allowing teams to innovate while adhering to financial constraints. Automated policy enforcement ensures that these guardrails are followed consistently, reducing the reliance on manual oversight and reducing the risk of human error. As a result, organisations can scale their cloud infrastructure confidently, knowing that they have the right governance in place to control costs and adhere to FinOps principles effectively.


## Policy Examples

Within the finops-stack, we've provided some policies we believe to be imported when considering implementing policy enforcement and FinOps practices, a few of these can be described below:


### Block Large images

Containers with excessively large image sizes can significantly impact the performance and efficiency of Kubernetes clusters. Large images take longer to pull from the registry, increasing deployment times and potentially causing delays in application start-up, which can disrupt critical operations. Additionally, these images consume more disk space on the nodes, which can lead to node resource exhaustion and reduce the available capacity for other workloads. Large images may also create unnecessary network traffic, particularly when pulled across multiple nodes, impacting overall cluster performance and increasing operational costs.

This policy is designed to mitigate such risks by enforcing a strict limit on container image sizes, ensuring that no container images larger than 2 Gibibytes are allowed to be deployed within the cluster. By blocking these images, the policy helps maintain efficient resource utilisation, prevents unintentional storage exhaustion, and reduces the likelihood of operational disruptions caused by oversized images. Furthermore, it acts as a safeguard against potential security risks where a malicious actor might intentionally deploy large images to disrupt cluster operations or cause denial of service (DoS) through resource exhaustion.

By enforcing this policy, organisations can maintain a standardised, efficient, and secure environment, ensuring that all workloads deployed are optimised for performance and resource usage. This policy also promotes best practices in image optimisation, encouraging development teams to reduce bloat, use multi-stage builds, and implement efficient base images to keep container sizes within acceptable limits.

### Validate Cost Center Label

Labels in Kubernetes are key/value pairs used to organise and manage workloads by adding metadata that is meaningful to users but does not affect the core functionality of the system. They play a crucial role in filtering, grouping, and selecting objects, such as pods, for operational tasks like deployments, scaling, monitoring, and cost tracking. Proper labelling ensures clarity and consistency across teams, enabling effective resource allocation, automation, and management within the cluster.

This policy enforces the requirement that all Pods deployed within the cluster must specify a label cost-center-label. The cost-center-label is essential for financial and operational governance, as it allows organisations to associate workloads with specific cost centres, departments, or projects. This enables accurate tracking of resource consumption, facilitating better financial oversight and cost allocation practices, which are particularly important in FinOps strategies.

By enforcing this policy, organisations ensure that all workloads are tagged with relevant financial metadata, which can be used to generate detailed cost reports and dashboards, monitor spending per team or project, and prevent unaccounted-for resource usage. It also promotes accountability and transparency across teams, ensuring that cloud resources are aligned with business objectives. Failure to apply this label may result in unidentified costs, making it difficult to optimise budgets or identify areas where cost-saving measures are needed.

### Restrict Scale of Deployments/Statefulsets, etc

Pod controllers, such as Deployments, manage replicas of pods to ensure high availability, scalability, and resiliency of applications running in Kubernetes. These controllers use the /scale subresource to dynamically adjust the number of pod replicas, either manually or through auto-scaling mechanisms. While scaling is critical to meeting demand, uncontrolled or excessive scaling can exhaust cluster resources, destabilise workloads, or lead to unexpected costs.

This policy, available starting in Kyverno 1.9, provides a collection of rules designed to enforce limits on the replica count in two scenarios: the initial creation of a Deployment and any subsequent scaling operations. By applying these rules, organisations can prevent deployments from creating or scaling to an excessively high or low number of replicas, which could disrupt cluster operations. Limiting replica counts ensures that resources are used efficiently and that scaling actions do not unintentionally degrade performance or overload the cluster.

The policy also accounts for operational governance by controlling the scale action directly through the /scale subresource. This ensures that replica limits are respected during runtime scaling operations, not just at deployment creation, giving administrators finer control over resource utilisation. This can help mitigate risks like uncontrolled scaling due to misconfigured horizontal pod autoscalers (HPA) or manual adjustments that exceed resource availability.

By enforcing this policy, organisations can optimise cluster stability, prevent resource exhaustion, and maintain cost-effective scaling practices, ensuring that both system performance and financial controls are upheld.

### Disallow Service LoadBalancers

In Kubernetes, the LoadBalancer service type provides a mechanism to expose services externally to the internet by automatically provisioning an external load balancer. While this is a convenient method to provide access to applications, improper or excessive use of LoadBalancer services can lead to unintended consequences, such as increased cloud costs, security risks, and resource consumption.

This policy enforces governance around the creation and usage of services with the type LoadBalancer, ensuring that they are only deployed under appropriate circumstances. By limiting or restricting the use of LoadBalancer services, the policy helps mitigate the risk of unexpected cloud costs associated with provisioning and maintaining external load balancers, which can accumulate rapidly, especially in large-scale environments. Additionally, it reduces exposure to potential security vulnerabilities by preventing the unnecessary exposure of internal services to the public internet.

By applying this policy, organisations can maintain tighter control over network traffic management, ensuring that only approved applications can utilise external load balancers, and promoting the use of alternative, cost-effective service types such as ClusterIP or NodePort where external access is not required. The policy also supports FinOps principles by enforcing the responsible use of cloud resources and minimising operational costs related to networking services.

### Add HPA to Deployments/StatefulSets

Horizontal Pod Autoscalers (HPA) are critical components in Kubernetes for maintaining application performance and resource efficiency by automatically adjusting the number of pod replicas in response to varying workloads. A Kyverno policy can automate the generation of HPA resources for each deployment, ensuring that applications are able to handle sudden spikes in traffic without downtime, while also scaling down during periods of low demand to conserve resources.

This policy dynamically creates HPA resources based on predefined scaling thresholds, such as CPU or memory usage, for every eligible deployment. By doing so, it provides a seamless mechanism for scaling applications according to real-time traffic patterns, preventing over-provisioning and ensuring that application performance is maintained during peak loads. In environments where traffic can be unpredictable, this proactive scaling approach minimises the risk of performance degradation, application downtime, or user dissatisfaction.

Additionally, the policy supports cost-efficiency by ensuring that resources are scaled down when demand decreases, avoiding the wasteful use of compute and memory resources. This helps organisations align their infrastructure usage with actual demand, supporting both operational and financial objectives, such as lowering cloud costs and optimising cluster performance. By automating the generation of HPA resources, the policy also reduces the operational burden on teams, ensuring that scaling is consistently and efficiently managed across the entire cluster without requiring manual intervention.

This Kyverno policy plays a crucial role in balancing application reliability, performance, and cost management, especially in dynamic cloud environments where traffic and resource requirements can fluctuate unpredictably.

### Prevent Naked Pods

Pods that are not managed by higher-level workload controllers, such as Deployments, StatefulSets, or DaemonSets, lack the self-healing and scaling capabilities that Kubernetes offers. These “naked” Pods, when deployed directly, are unsuitable for production environments as they do not benefit from automated recovery if a node fails or a pod crashes, nor can they scale up or down based on demand.

This policy enforces a restriction that prevents the creation of such standalone Pods unless they are part of a workload controlled by higher-level Kubernetes objects. By doing so, the policy ensures that all applications deployed within the cluster benefit from the inherent resilience, scalability, and manageability provided by workload controllers. This is crucial in production environments where reliability and the ability to automatically recover from failures are necessary to maintain application uptime and service availability.

In addition to improving operational stability, the policy promotes best practices in Kubernetes resource management by encouraging developers and operators to use controllers like Deployments for managing pod lifecycles. Workload controllers provide declarative updates, rolling deployments, and the ability to maintain desired state configurations, which are essential features for maintaining robust and scalable applications in modern cloud-native environments.

By enforcing this policy, organisations can prevent the accidental deployment of fragile, unmanaged Pods, reducing operational risks and ensuring a more resilient production environment.
