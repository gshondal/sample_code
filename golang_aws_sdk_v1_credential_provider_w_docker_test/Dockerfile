#Dockerfile

#Go Version being used
FROM golang

#Environmental variables
ENV AWS_REGION=us-east-1
ENV HOMEDIR=/opt/golang
ENV GO111MODULES=on

#Creating Home directory
RUN mkdir -p $HOMEDIR
WORKDIR $HOMEDIR

#Getting AWS SDK
RUN go mod init proj
RUN go get github.com/aws/aws-sdk-go
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go//service/sts

#Deploying code to image
COPY src .
RUN go build sample.go

CMD ./sample