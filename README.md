# scandns
Scanner for DNS open resolvers

* load your CIDRs in `cidr_to_scan.json`
* `make run`

Example output:
```
2019/04/30 13:55:23 ---- [RESULTS] ----
116.X.2.96
116.X.3.116
116.X.16.44
116.X.20.195
116.X.111.142
116.X.112.219
2019/04/30 13:55:23 ---- [END] ----
2019/04/30 13:55:23 ---------------------------------------------------------
2019/04/30 13:55:23 DONE: (5842) IP addresses in (46.632074735s) and found (6) possible candidates.
2019/04/30 13:55:23 ---------------------------------------------------------
```