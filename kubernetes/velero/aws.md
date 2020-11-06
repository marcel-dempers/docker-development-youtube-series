# Run AWS CLI

```
docker run -it --rm --entrypoint /bin/sh amazon/aws-cli:2.0.55

# install JSON tool
yum install -y jq
```

# Login to AWS

https://docs.aws.amazon.com/eks/latest/userguide/getting-started-console.html

```
# Access your "My Security Credentials" section in your profile. 
# Create an access key

aws configure

Default region name: ap-southeast-2
Default output format: json


```

# Create a Storage Bucket

```
BUCKET=veleromarcel
REGION=ap-southeast-2
aws s3api create-bucket --bucket $BUCKET --region $REGION --create-bucket-configuration LocationConstraint=$REGION
```

# Create IAM User

```
aws iam create-user --user-name velero
```

# Setup Policy for the User

```
cat > velero-policy.json <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeVolumes",
                "ec2:DescribeSnapshots",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:CreateSnapshot",
                "ec2:DeleteSnapshot"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:DeleteObject",
                "s3:PutObject",
                "s3:AbortMultipartUpload",
                "s3:ListMultipartUploadParts"
            ],
            "Resource": [
                "arn:aws:s3:::${BUCKET}/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:ListBucket"
            ],
            "Resource": [
                "arn:aws:s3:::${BUCKET}"
            ]
        }
    ]
}
EOF

aws iam put-user-policy \
  --user-name velero \
  --policy-name velero \
  --policy-document file://velero-policy.json

```

# Create Access Key for that user

```
aws iam create-access-key --user-name velero > /tmp/key.json

AWS_ACCESS_ID=`cat /tmp/key.json | jq .AccessKey.AccessKeyId | sed s/\"//g`
AWS_ACCESS_KEY=`cat /tmp/key.json | jq .AccessKey.SecretAccessKey | sed s/\"//g`

```

# Export variables

Let's export these variables into our Velero container <br/>
<br/>
Copy and paste this to the velero container:
```

printf "export AWS_ACCESS_ID=$AWS_ACCESS_ID \nexport AWS_ACCESS_KEY=$AWS_ACCESS_KEY\nexport BUCKET=$BUCKET \nexport REGION=$REGION\n"
```