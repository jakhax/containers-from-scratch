# Containers From Scratch 
- While learning about [Os level virtualization](https://en.wikipedia.org/wiki/OS-level_virtualization) to create a sandbox environment to execute untrusted user code,
I came across the concept of containers.
- I had used containers a couple of times, specifically docker containers, before but only understood it from a very general perspective, 
`It isolates an application and its dependencies into a self-contained unit that can run anywhere.`
- In my opinion, containers are really trending/hyped technology that is poorly understood, so I think it would be nice to create one by yourself to understand how it works.

## Running this example 
- You'll need to do it from a linux machine.
- First download the ubuntu file system, you can get it from here, [ubuntu_fs.zip](https://drive.google.com/file/d/1Qunrj-3_oXhPQCz4aqaYffiGBY9DdRR9/view?usp=sharing),  unzip it at the root of this project.

```bash
# needs root privilege for creating cgroup
sudo su
go run main.go run /bin/bash
``` 

### How it works

- Parent process fork itself with `CLONE_NEWUTS`, `CLONE_NEWPID`, `CLONE_NEWNS` flags with isolated hostname, processes and mounts
- The forked process will create `cgroup` to limit memory usage of itself and any child process it creates
- Mount `./ubuntu_fs` directory as root filesystem using `chroot` to limit access to host machine's filesystem
- Mount `/mytemp` directory as tmpfs. Any change made to this directory will not be visible from host.
- Mount proc (where `CLONE_NEWPID` namespace was already set) so that container can run `ps` and see only the processes running inside it.
- Execute the supplied argument `/bin/bash` inside the isolated environment


## Resources
How this work is already well explained by the following resources/tutorials
- [Containers From Scratch by Liz Rice (video)](https://www.youtube.com/watch?v=8fi7uSYlOdc)
- [Build your own container from Scratch](https://www.infoq.com/articles/build-a-container-golang/)
- [Docker From Scratch with python workshop](https://github.com/Fewbytes/rubber-docker)
- [A Beginner-Friendly Introduction to Containers, VMs and Docker](https://www.freecodecamp.org/news/a-beginner-friendly-introduction-to-containers-vms-and-docker-79a9e3e119b/)
- [Namespaces in Go Series](https://medium.com/@teddyking/linux-namespaces-850489d3ccf)
- [Understand Container](https://pierrchen.blogspot.com/2018/08/understand-container-index.html)

## Security of Containers
 **Containers are not contained**
 - Containers by default are not a secure environment to execute untrusted code(e.g Saas), an application running in your container can still exploit a vulnerability in the linux kernel.
 - For ways to secure your containers you can take a look at [gVisor](https://github.com/google/gvisor),  [Kata Containers](https://katacontainers.io/)
 and [Firecracker](https://aws.amazon.com/blogs/aws/firecracker-lightweight-virtualization-for-serverless-computing/)
 
 ### Resources on security of Containers
 - [Sandboxing your Containers with gVisor (video)](https://www.youtube.com/watch?v=kxUZ4lVFuVo)
 - [Is docker as a sandbox secure](https://security.stackexchange.com/questions/107850/docker-as-a-sandbox-for-untrusted-code)
 - [Open Sourcing gVisor](https://cloud.google.com/blog/products/gcp/open-sourcing-gvisor-a-sandboxed-container-runtime)
 - [Enemy within; Running untrusted code with gVisor (video)](https://www.youtube.com/watch?v=1Ib-rfSzDuM)
