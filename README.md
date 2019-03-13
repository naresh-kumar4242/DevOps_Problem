Maintained Idempotency of Automation

- We can run AlphaServer anywhere (Bare metal,VM or container(Single container or inside Kubernetes pod)), but AlphaClient need to be run in VM to collect login attempts.
  - Although AlphaClient can also be run in a container on a VM after mapping/mounting host's /var/log/ inside container
  - This will allow us to access host's auth.log file , But we will run it on VM, keeping things simple
- Please clone the repo.
- We will run AlphaServer in Docker container, exposing port 9001
  - cd DevOps_Problem/
  - docker build -t server:v1 ./AlphaServer/
  - docker run -it -p 9001:9001 server:v1
- Now that our AlphaServer is running, We will run AlphaClient in a VM . VM can be provisioned locally using Vagrant or in cloud such as AWS EC2 instance.
- We just need to change server address in AlphaClient code or we can get it through ENV variable
  - We can also get attempts on particular date after changing Date inside AlphaClient code
- We can also directly run on our own host to get all login attempts.
- Simply , Run
  - cd AlphaClient/
  - go run client.go
- Please open following in browser :
  - http://localhost:9001/ssh_details
  - You should be able to see all login attempts along with Host/Node name as well as particular date


- We can deploy as many replicas of AlphaServer as we need inside Kubernetes cluster. AlphaClient can be deployed to as many VMs as we want using tools such as Ansible,Terraform,Chef.
- I covered the most basic scenario of getting login attempts inside our own system
