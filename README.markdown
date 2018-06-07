# Hi Candidate!

## Description of main task
We've implemented this simple transaction service, it's able to add, get, list and delete transactions.
After playing around with the transaction service we discovered that we really wanted to know the balance
of given accounts over time, so we created a protobuf specification for a new service. This balance service
may serve as a datasource for a graph, plotting the balance of an account over time. We hope you can help us
implement the service and some CLI for querying the service.

We envisioned a CLI looking something like this:

Here we query to balance service to return balance for `accounts/1`, period
is specified and the size of each bucket (720h, 30 days) is specified. We expect that the service
would figure out when these buckets start and stop and also what the balance is at that given period.

```bash
$ bin/candidate balance --duration 720h accounts/1 2016-01-01 2016-01-31
from: 2016-01-01 to: 2016-02-01 amount: 1234
from: 2016-02-01 to: 2016-03-01 amount: 4321
from: 2016-03-01 to: 2016-04-01 amount: -1234.5
...
from: 2016-12-01 to: 2016-12-31 amount: 1
```

Another way to query the balance service should be to specify the number of buckets we want,
for example:

```bash
$ bin/candidate balance --buckets 365 accounts/1 2016-01-01 2016-12-31
from: 2016-01-01 to: 2016-01-02 amount: 123
...
from: 2016-12-30 to: 2016-12-31 amount: 1
```

This command should probably return the daily balance of that account.

The specification for the balance service is defined in `api/balance.proto`, feel free to modify
it if you come up with something better. An empty service implementation is added in `pkg/balance/service.go`, add meat
and bones to it.

## Tasks

1. Implement the missing balance service
2. Add required support in CLI (for getting balance results)
3. Add tests for newly added functionality

## Bonus tasks

- Add better parameter support in CLI (listen address, transaction endpoint etc)
- Implement balance service or client in another language
- Improve the non-existing test coverage
- Improve error handling
- Add HTTP REST endpoints for balance service
- Add REST gateway to service launch code
- Add kubernetes deployment configuration

# Be creative, feel free

Looking forward for your submissions,

Sincerely, SMB LAB geeks

(PS: You can tarball your solution and ship it to Torkel <torkel at smblab.no>)
