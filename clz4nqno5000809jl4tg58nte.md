---
title: "Understanding Amazon Elastic Block Store (EBS): An Introduction"
seoTitle: "Intro to Amazon Elastic Block Store"
seoDescription: "Introduction to Amazon EBS: learn about types, attachment, and uses of SSD and HDD-backed volumes for EC2 instances"
datePublished: Sat Jul 27 2024 21:42:47 GMT+0000 (Coordinated Universal Time)
cuid: clz4nqno5000809jl4tg58nte
slug: understanding-amazon-elastic-block-store-ebs-an-introduction
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722116526404/8002ccf6-14e9-4858-bde9-6872808c2159.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722116554803/0c650e31-8d35-4596-ad98-ca1d07d7d1d1.png
tags: aws, ebs, aws-ebs, aws-ebs-introduction, aws-ebs-beginners, aws-ebs-for-beginners-introduction, ebs-volume-types

---

# **Introduction**

1. Amazon EBS lets you create and attach storage volumes to EC2 instances.
    
2. You can attach multiple EBS volumes to an instance.
    
3. Once the storage volume is created, you can set up a file system on it. Then, you can run a database, store files and applications, or use it as a block device in other ways.
    
4. EBS volumes must be in the same availability zone as the instances they are attached to.
    
5. Root EBS volumes are deleted on termination by default.
    
6. Amazon EBS volumes are automatically replicated to protect you from the failure of a single component.
    
7. You can also attach additional EBS volumes to a running instance.
    
8. EBS volume attached to the EC2 instance where Windows or Linux is installed is known as the Root device of volume.
    

# **Type Of Amazon EBS Volume**

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722109769022/1652f7f5-57d4-49bd-a0e8-fb0281688600.png align="center")

Amazon EBS provides two types of volume that differ in performance characteristics and price. EBS Volume types fall into two parts:

1. SSD-Backed Volumes
    
2. HDD-Backed Volumes
    

## **SSD-Backed Volumes**

1. SSD stands for solid-state Drives.
    
2. In June 2014, SSD storage was introduced.
    
3. It is a general-purpose storage.
    
4. Can be a boot volume.
    
5. SSD storage is very high-performing, but it is quite expensive as compared to HDD (Hard Disk Drive) storage.
    

**SSD is further classified into two parts:**

1. General Purpose SSD
    
2. Provisioned IOPS SSD
    

## **HDD-Backed Volumes**

1. It stands for Hard Disk Drive
    
2. HDD-based storage was introduced in 2008.
    
3. The size of the HDD based storage could be between 125 GiB to 1 TiB.
    
4. Cannot be a boot volume.
    
5. EBS Multi-attach not supported.
    

**HDD is further classified into three parts:**

1. Throughput Optimized HDD (ST1)
    
2. Cold HDD (SC1)
    
3. Magnetic Volume
    

For more information do check out [Amazon EBS](https://aws.amazon.com/ebs/).