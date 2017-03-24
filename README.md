# carbon-schema-to-clickhouse
Converts storage-schema.conf from graphite/carbon to clickhouse graphite_rollup XML

Migrating our `storage-schemas.conf` from an older graphite implementation to
the XML snippet needed for [ClickHouse](https://github.com/yandex/ClickHouse)
would have needed much more time than hacking this small piece of go code.

## Usage:

```
% cat storage-schemas-example.conf                                                                                                          [carbon]
pattern = ^carbon\.
retentions = 60:90d,60m:120d,360m:400d,720m:2y,1d:3y,7d:10y

[collectd]
pattern = ^collectd\..*
retentions = 10s:12h,1m:10d,15m:20d,30m:40d,60m:120d,360m:400d,720m:2y,1d:3y,7d:10y

[netapp.capacity]
pattern = ^netapp\.capacity\.*
retentions = 15m:100d, 1d:5y

[default]
pattern = .*
retentions = 1m:10d,15m:20d,30m:40d,60m:120d,360m:400d,720m:2y,1d:3y,7d:10y
```

```
% ./carbon-schema-to-clickhouse -h
Usage of ./carbon-schema-to-clickhouse:                                                                 
  -rollupfunction string
    	clickhouse rollup function (default "any")
  -schemafile string
    	storage schema file to convert (default "/etc/carbon/storage-schemas.conf")
```
```
% ./carbon-schema-to-clickhouse -schemafile storage-schemas-example.conf -rollupfunction avg
<graphite_rollup>
	<!-- carbon -->
	<pattern>
		<regexp>^carbon\.</regexp>
		<function>avg</function>
		<retention>
			<age>0</age>
			<precision>60</precision>
		</retention>
		<retention>
			<age>7776000</age>
			<precision>3600</precision>
		</retention>
		<retention>
			<age>10368000</age>
			<precision>21600</precision>
		</retention>
		<retention>
			<age>34560000</age>
			<precision>43200</precision>
		</retention>
		<retention>
			<age>63072000</age>
			<precision>86400</precision>
		</retention>
		<retention>
			<age>94608000</age>
			<precision>604800</precision>
		</retention>
	</pattern>
	<!-- collectd -->
	<pattern>
		<regexp>^collectd\..*</regexp>
		<function>avg</function>
		<retention>
			<age>0</age>
			<precision>10</precision>
		</retention>
		<retention>
			<age>43200</age>
			<precision>60</precision>
		</retention>
		<retention>
			<age>864000</age>
			<precision>900</precision>
		</retention>
		<retention>
			<age>1728000</age>
			<precision>1800</precision>
		</retention>
		<retention>
			<age>3456000</age>
			<precision>3600</precision>
		</retention>
		<retention>
			<age>10368000</age>
			<precision>21600</precision>
		</retention>
		<retention>
			<age>34560000</age>
			<precision>43200</precision>
		</retention>
		<retention>
			<age>63072000</age>
			<precision>86400</precision>
		</retention>
		<retention>
			<age>94608000</age>
			<precision>604800</precision>
		</retention>
	</pattern>
	<!-- netapp.capacity -->
	<pattern>
		<regexp>^netapp\.capacity\.*</regexp>
		<function>avg</function>
		<retention>
			<age>0</age>
			<precision>900</precision>
		</retention>
		<retention>
			<age>8640000</age>
			<precision>86400</precision>
		</retention>
	</pattern>
	<!-- default -->
	<default>
		<function>avg</function>
		<retention>
			<age>0</age>
			<precision>60</precision>
		</retention>
		<retention>
			<age>864000</age>
			<precision>900</precision>
		</retention>
		<retention>
			<age>1728000</age>
			<precision>1800</precision>
		</retention>
		<retention>
			<age>3456000</age>
			<precision>3600</precision>
		</retention>
		<retention>
			<age>10368000</age>
			<precision>21600</precision>
		</retention>
		<retention>
			<age>34560000</age>
			<precision>43200</precision>
		</retention>
		<retention>
			<age>63072000</age>
			<precision>86400</precision>
		</retention>
		<retention>
			<age>94608000</age>
			<precision>604800</precision>
		</retention>
	</default>
</graphite_rollup>

```

## Thanks
Thanks to Roman Lomonosov for his [go-carbon](https://github.com/lomik/go-carbon) code - I'm reusing the storage schema parser here.
