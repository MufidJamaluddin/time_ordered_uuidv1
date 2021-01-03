# Ordered UUID V1

![Github CI](https://github.com/MufidJamaluddin/uuidv1_ordered/workflows/Go/badge.svg)

Ordering UUID V1 bytes by inserted timestamp for practical purposes. If you want to store UUID's sequentially in databases, you can use this library.

## Requirement

Go 1.14 ~ 1.15, other versions are not tested

## Instalation

`go get github.com/MufidJamaluddin/uuidv1_orderer`

## How to Use

1. Ordering Existing UUID's

```
unorderedUuid := UUIDStringToBytes("c5db1800-ce4c-11de-a5e2-1b45123c6e98")

orderedUuid := ToOrderedUuid(unorderedUuid)

orderedUuidStr := UUIDBytesToString(orderedUuid)

// next, your logic
```

2. Transform Back to Standard UUID V1

```
orderedUuid := UUIDStringToBytes("c5db1800-ce4c-11de-a5e2-1b45123c6e98")

standardUuidV1 := FromOrderedUuid(orderedUuid)

standardUuidV1Str := UUIDBytesToString(standardUuidV1)

// next, your logic
```

## Notes

The differences explanation between ordered and standard UUID V1 is bellow

### UUID V1 In String

Example: `c5db1800-ce4c-11de-a5e2-1b45123c6e98`

Meaning: 
1. Timestamp `1de-ce4c-c5db1800` in `c5db1800-ce4c-11de`
2. Standard UUID version `1` in `c5db1800-ce4c-(1)1de`
3. Clock sequence `a5e2`
4. Standard MAC Address or Nonstandard Random IDs `1b45123c6e98`

### Ordered UUID V2

Reorder timestamp to it's origin form with UUID version on top.

Example: `11dece4c-c5db-1800-a5e2-1b45123c6e98`

## Licenses

(c) Mufid Jamaluddin. MIT Licenses
