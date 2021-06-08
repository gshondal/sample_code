# Build image
`docker build . -t golang_sample`

# Run and Mount Credentials Directory
# Linux: 
`sudo docker run -v ~/.aws:/root/.aws:ro golang_sample`

# Windows:
`docker run -v c:/Users/Username/.aws:/root/.aws:ro golang_sample`

# Output
```
{
  Account: "0123456789",
  Arn: "arn:aws:iam::0123456789:user/iamUser",
  UserId: "AIDA123P456D789"
}
```
