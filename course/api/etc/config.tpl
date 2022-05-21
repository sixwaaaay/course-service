# touch a file named course.yaml and fill it like this
Name: course
Host: {host}
Port: {port}
CourseRpc:
  Etcd:
    Hosts:
      - {host1}
    Key: course.rpc
SchoolRpc:
  Etcd:
    Hosts:
      - {host1}
    Key: school.rpc
MinIO:
  Endpoint: {host}:{port}
  AccessKey: {accessKey}
  SecretKey: {secretKey}
  Bucket: {bucket}
