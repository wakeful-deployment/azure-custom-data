# `azure-custom-data`

## What is this for?

When creating a VM in Azure, one can provide custom data. The data is
base64 encoded and buried in an xml file. This is a small binary that
will fine, parse, and output the string.

If you put json in the string it becomes even more magical.

## An example

During VM creation, assume I posted
`eyJuYW1lIjoiZXhhbXBsZSIsInNpemUiOiJCYXNpY19BMSIsImxvY2F0aW9uIjoiZWFzdHVzIiwicmVzb3VyY2VfZ3JvdXAiOiJleGFtcGxlIiwiY3JlYXRlZF9hdCI6IjE0NTE0ODY4MzIuNzIwMjMwOCJ9`
in the `customData` field in `osProfile` in `properties` for the vm.

```sh
$ sudo azure-custom-data
{"name":"example","size":"Basic_A1","location":"eastus","resource_group":"example","created_at":"1451486832.7202308"}
```

:boom:
