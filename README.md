I forked [this repository](https://github.com/guce/ec2-start-stop) and I added start&stop RDS golang file.


# EC2&RDS Start Stop

Implement EC2&RDS scheduled start and stop with **Lambda golang function** & **CloudWatch**.



# Step

## IAM

### New Policy
#### For EC2
```shell
IAM -> Policies -> Create policy

Visual editor
Service -> EC2
Actions -> Access level -> Write -> StartInstances,StopInstances
Resources -> All resources

PolicyName -> StartStopEC2Instances
```
#### For RDS
```shell
IAM -> Policies -> Create policy

Visual editor
Service -> RDS
Actions -> Access level -> Write -> StartDBInstance,StopDBInstance
Resources -> All resources

PolicyName -> StartStopRDSInstances
```



### New Role
#### For EC2
```shell
IAM -> Roles -> Create role

AWS service -> Lambda
Policy -> StartStopEC2Instances

Role name -> StartStopEC2Instances
```
#### For RDS
```shell
IAM -> Roles -> Create role

AWS service -> Lambda
Policy -> StartStopRDSInstances

Role name -> StartStopRDSInstances
```

## Lambda Function

```shell
AWS Lambda -> Functions -> Create function

Name -> StartEC2Instances
Runtime -> Go 1.x
Role -> Existing role -> StartStopEC2Instances
```

```shell
clone this repository
./build.sh
```

```shell
Function code
Function package -> start-ec2-instances.zip
Handler -> start-ec2-instances

Test -> Create new test event
{
  "InstanceRegion": "ap-northeast-2",
  "InstanceIdList":[
    "i-xxxxxxxxxxxxxxxxx",
    "i-xxxxxxxxxxxxxxxxx"
  ]
}

if its a RDS
{
  "InstanceRegion": "ap-northeast-2",
  "InstanceIdList":[
    "{DB_identifier_name}",
    "{DB_identifier_name2}"
  ]
}

```



## CloudWatch

```shell
Rules -> Create role -> Schedule

Cron expression -> xxxxx
#e.g.
00 00 ? * MON-FRI *		#9 am(GMT+9) every working day, start ec2 instances
00 10 ? * * *			#19 pm(GMT+9) every day, stop ec2 instances

Targets -> Lambda function -> StartEC2Instances
Constant(JSON text)
{
  "InstanceRegion": "ap-northeast-2",
  "InstanceIdList":[
    "i-xxxxxxxxxxxxxxxxx",
    "i-xxxxxxxxxxxxxxxxx"
  ]
}

if its a RDS
{
  "InstanceRegion": "ap-northeast-2",
  "InstanceIdList":[
    "{DB_identifier_name}",
    "{DB_identifier_name2}"
  ]
}
```



# Reference

AWS lambda go  https://github.com/aws/aws-lambda-go

AWS sdk go https://github.com/aws/aws-sdk-go

CloudWatch https://docs.amazonaws.cn/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
