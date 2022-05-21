# touch a file named course.yaml and fill it like this
Name: course.rpc
ListenOn: {host}:{port}
Etcd:
  Hosts:
  - {host1}
  Key: course.rpc
DB:
  DataSource: {dsn}
