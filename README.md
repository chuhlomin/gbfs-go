# gbfs-go

## What is GBFS

[The General Bikeshare Feed Specification][], known as GBFS, is the open data standard for bikeshare. GBFS makes real-time data feeds in a uniform format publicly available online, with an emphasis on findability. GBFS is intended to make information publicly available online; therefore information that is personally identifiable is not currently and will not become part of the core specification.

## What is gbfs-go

gbfs-go is a [Go][] library that aimed to simplify reading and writing GBFS data, it provides Go structures and HTTP-client wrapper.

## Examples

* systems
* gbfs
* systems_information
* station_information
* station_status
* free_bike_status
* systems_hours

Try examples:

```bash
go run examples/system_information/main.go
```

## License

MIT

Keep in mind that data feeds may have a [different license] (like CC0-1.0, CC-BY-4.0, etc).

[The General Bikeshare Feed Specification]: https://github.com/NABSA/gbfs
[Go]: https://golang.org
[different license]: https://github.com/NABSA/gbfs/blob/master/data-licenses.md
