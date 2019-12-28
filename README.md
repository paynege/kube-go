# kube-go
## Preparation

## Process
1. Create installation directory
    * etcd
        ```
        mkdir Config.App["etcd"].AppTargetPath/{bin,cfg,ssl}
        ```
    
    * kubernetes
        ```
        mkdir Config.App["kubernetes"].AppTargetPath/{bin,cfg,ssl}
        ```
2. Generate cert files
    * json files according to template
        ```
        ETCD_CA = "etcd_ca/"
        ```
    * cert files
3. Generate conf files
    * etcd
    * kubernetes
4. Generate startup service files
    * etcd
    * kubernetes
5. Copy binary files to installation directory
    * etcd
    * kubernetes
6. Run Apps
    * etcd
    * kubernetes
        1. kube-apiserver
        2. kube-controller-manager
        3. kube-scheduler
        4. kubelet
        5. kube-proxy