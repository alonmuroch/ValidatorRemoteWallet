# Patchs for dependencies

----
## What is this?
see [here](https://github.com/prysmaticlabs/prysm/tree/master/third_party)

----
## Why it's here?
The remote wallet uses [ethereumapi](https://github.com/prysmaticlabs/ethereumapis) as a dependency, enabling the wallet to have its own slashing logic.

The way prys uses it is to serialize different data structures that are defined at [ethereumapi](https://github.com/prysmaticlabs/ethereumapis).

Those data structures have different tags for serializing them, issue is, those tags are added in the ethereumapi patch prysm implemented 

----
## When to update it?
Whenever prysm updates the [ethereumapi](https://github.com/prysmaticlabs/ethereumapis) patch, we need to do so as well.
