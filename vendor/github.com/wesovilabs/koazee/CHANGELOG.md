# Changelog

## [Unreleased]

### Added

- Operation reverse


## Gelada (v0.0.2) - 1/12/2018

### Added
- Benchmark testing for all the operations [Koazee Benchmark Report](https://github.com/wesovilabs/koazee/wiki/Benchmark-Report)
- Working with generated code instead of reflection for primitive streams
- Several changes in code to get a better performance
- Full wiki [Koazee wiki](https://github.com/wesovilabs/koazee/wiki)
- Caching validation types in operations
- New examples can be found [here](https://github.com/wesovilabs/koazee/tree/master/samples)
### Removed
- Compose operation
- Interface S
- Logger is deprecated
- External Site has been removed


## Titi (v0.0.1) - 12/11/2018
### Added
- add: Add a new element into the stream.
- at: Obtain the element in the stream that is in the given position
- compose: Join 2 or more streams in a single one
- contains: Check if an element is found in the stream
- count: Return the number of elements in the stream
- drop: Drop an existing element in the stream
- filter: Discard those elements in the stream that do not match with the given conditions
- first: Obtain the first element in the stream
- foreach: Do something over all the elements in the stream
- last: Obtain the last element in the stream
- map: Convert the current elements in the stream into a different type
- reduce: Return the result after applying the provided function over all the items in the stream
- removeduplicates: Remove duplicates elements in the stream
- sort: Sort the elements in the stream
