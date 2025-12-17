provider "aws" {
    region = "us-west-2"
}
resource "aws_instance" "ocp_node" {
    ami = "ami-12345678"
    instance_type = "m5.large"
}