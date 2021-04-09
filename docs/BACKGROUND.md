# Background

## Deployment Frequency KPI

### References

- [Accelerate - The Science Behind DevOps: Building and Scaling High Performing Technology Organizations
By Nicole Forsgren, Jez Humble, Gene Kim · 2018](https://itrevolution.com/book/accelerate/)

    From [State of DevOps 2019](https://services.google.com/fh/files/misc/state-of-devops-2019.pdf):

    "*Deployment frequency - The elite group reported that it routinely deploys on-demand and performs multiple deployments per day, consistent with the last several years. By comparison, low performers reported deploying between once per month (12 per year) and once per six months (two per year), which is a decrease in performance from last year. The normalized annual deployment numbers range from 1,460 deploys per year **(calculated as four deploys per day x 365 days) for the highest performers** to seven deploys per year for low performers (average of 12 deploys and two deploys). Extending this analysis shows that elite performers deploy code 208 times more frequently than low performers. It's worth noting that four deploys per day is a conservative estimate when comparing against companies such as CapitalOne that report deploying up to 50 times per day for a product, **8 for companies such as Amazon, Google, and Netflix** that deploy thousands of times per day (aggregated over the hundreds of services that comprise their production environments).*"

- [CDF Member Webcast: Data-Driven Benchmarks for High Performing Engineering Teams](https://www.youtube.com/watch?v=iUFpRFvlT2U):

    Deployment Frequency ("Throughput") per day per project:

    | Percentile | 2020 Value | 2019 Value |
    | --- | --- | --- |
    | 5p | 0.03 | 0.03 |
    | 50p | 0.70 | 0.80 |
    | 90p | 16.03 | 13.00 |
    | 95p | 32.125 | 25.47 |
    | **Mean** | **8.22** | **5.76** |

=> This is where the [Grafana Dashboard Gauges'](../grafana-dashboards/) thresholds are coming from.

## Unnamed

- https://thenewstack.io/how-devops-affects-business-stakeholders-and-leaders/
- https://www.thundra.io/apm
- https://amazicworld.com/dealing-with-devops-metrics-and-kpis/
- https://www.cloudbees.com/blog/engineering-teams-health
- https://www.cloudbees.com/products/engineering-efficiency
