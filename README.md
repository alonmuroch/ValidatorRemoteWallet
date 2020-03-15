# Prysm Validator Remote Wallet

----
## What is this?
This is a remote wallet manager for a prysm validator. 
The idea is to be able to connect to a validator infrastructure (might not be controlled by you) and still maintain a high level of security.

a.k.a hold and manage your own keys 

----
## Building
1. [protobuf](https://github.com/golang/protobuf) on commit 6c65a5562fc06764971b7c5d05c76c75e84bdbf7. This is because prysm uses an older version of grpc which is not supported by newer protoc gen versions
2. build the Proto files to a go target
3. run bazel run //:gazelle to rebuild the Build files

----
