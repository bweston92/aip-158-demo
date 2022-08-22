# AIP 158 Demo

**WARNING! This is self documenting my first attempt at implementing the AIP-158 and could be seriously wrong.**

Blog post: https://bweston.me/a-grpc-pagination-api-implementation-aip-158-c74014d7be24

## Components

### svc-contacts

This is a small API that implements a service named `Contacts` and has one method `ListContacts`
where `ListContactsRequest` and `ListContactsResponse` implement AIP-158 for pagination.

### app-example-1

An example where the pagination is done not using a page token at all to go through the pages.

As described in the blog post, this means that if new results come in you will see some of the results
you've already seen on the following pages.

### app-example-2

An example where the pagination is done using a page token to go through the pages.

### app-example-3

An example where the pagination is done with a session state and has a custom implementation for the
actual frontend application.
