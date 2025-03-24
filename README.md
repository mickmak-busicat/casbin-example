# casbin-example

## Performance matrix rough numbers:

The performance depends on the order of the policy. Later policy need more time to be matched by the enforce function.

With domain:

|Policy rows|Init time|Enforce first policy|Enforce last policy|Estimated Total|
| -------- | ------- | -------- | ------- | ------- |
|5,000|15ms|200us|21ms|40ms|
|10,000|30ms|200us|60ms|90ms|
|20,000|50ms|200us|200ms|250ms|


Without domain:

|Policy rows|Init time|Enforce first policy|Enforce last policy|Estimated Total|
| -------- | ------- | -------- | ------- | ------- |
|5,000|15ms|200us|19ms|40ms|
|10,000|30ms|200us|60ms|90ms|
|20,000|50ms|200us|200ms|250ms|

### Result: performance with domain is similar to without domain.

## Policy row meaning:

Lets assume,
- one role contains 20 permissions
- one enforcer only check one user

10,000 rows = 10000/20 = ~500 roles on one user
