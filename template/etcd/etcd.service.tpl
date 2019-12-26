[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile={{ etcd-target-path }}/cfg/etcd.conf
ExecStart=/opt/etcd/bin/etcd \
--peer-cert-file={{ etcd-target-path }}/ssl/server.pem \
--peer-key-file={{ etcd-target-path}}/ssl/server-key.pem \
--peer-trusted-ca-file={{ etcd-target-path }}/ssl/ca.pem
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target