#!/bin/bash

protoc ./discountpb/discount.proto --go_out=plugins=grpc:.