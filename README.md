# Ordered UUID V1

Transform UUID V1 bytes to ordered UUID V1 bytes by inserted time for practical purposes. If you want to store UUID's sequentially in databases, you can use this library.

### Notes

1. UUID V1 In String

`c5db1800-ce4c-11de-a5e2-1b45123c6e98`

Meaning: 
1. Timestamp `1de-ce4c-c5db1800` in `c5db1800-ce4c-11de`
2. Standard UUID version `1` in `c5db1800-ce4c-(1)1de`
3. Clock sequence `a5e2`
4. Standard MAC Address or Nonstandard Random IDs `1b45123c6e98`

2. Ordered UUID V2

Reorder timestamp to its origin form with UUID version on top.