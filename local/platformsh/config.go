// Code generated by platformsh/generator/main.go
// DO NOT EDIT

/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package platformsh

var availablePHPExts = map[string][]string{
	"amqp":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"apc":             {"5.4", "5.5"},
	"apcu":            {"5.4", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"apcu_bc":         {"7.0", "7.1", "7.2", "7.3", "7.4"},
	"applepay":        {"7.0", "7.1", "7.4"},
	"bcmath":          {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"blackfire":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"bz2":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"calendar":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ctype":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"curl":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"dba":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"dom":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"enchant":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"event":           {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"exif":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ffi":             {"7.4", "8.0", "8.1"},
	"fileinfo":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ftp":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"gd":              {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"gearman":         {"5.4", "5.5", "5.6"},
	"geoip":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"gettext":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"gmp":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"http":            {"5.4", "5.5", "7.3", "7.4", "8.0", "8.1"},
	"iconv":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"igbinary":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"imagick":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"imap":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"interbase":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0"},
	"intl":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ioncube":         {"7.0", "7.1", "7.2"},
	"json":            {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ldap":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"mailparse":       {"7.0", "7.1", "7.2", "7.4", "8.0", "8.1"},
	"mbstring":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"mcrypt":          {"5.4", "5.5", "5.6", "7.0", "7.1"},
	"memcache":        {"5.4", "5.5", "5.6"},
	"memcached":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"mongo":           {"5.4", "5.5", "5.6"},
	"mongodb":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"msgpack":         {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0"},
	"mssql":           {"5.4", "5.5", "5.6"},
	"mysql":           {"5.4", "5.5", "5.6"},
	"mysqli":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"mysqlnd":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"newrelic":        {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"oauth":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"odbc":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"opcache":         {"5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo_dblib":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo_firebird":    {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"pdo_mysql":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo_odbc":        {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo_pgsql":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo_sqlite":      {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pdo_sqlsrv":      {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pgsql":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"phar":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pinba":           {"5.4", "5.5", "5.6"},
	"posix":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"propro":          {"5.6"},
	"pspell":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pthreads":        {"7.1"},
	"raphf":           {"5.6", "7.4", "8.0", "8.1"},
	"rdkafka":         {"8.1"},
	"readline":        {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"recode":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3"},
	"redis":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"shmop":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"simplexml":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"snmp":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"soap":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sockets":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sodium":          {"7.2", "7.3", "7.4", "8.0", "8.1"},
	"sourceguardian":  {"7.0", "7.1"},
	"spplus":          {"5.4", "5.5"},
	"sqlite3":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sqlsrv":          {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ssh2":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sybase":          {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sysvmsg":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sysvsem":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"sysvshm":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"tideways":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"tideways_xhprof": {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"tidy":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"tokenizer":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"uuid":            {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"wddx":            {"7.0", "7.1", "7.2", "7.3", "7.4"},
	"xdebug":          {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"xcache":          {"5.4", "5.5"},
	"xhprof":          {"5.4", "5.5", "5.6"},
	"xml":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"xmlreader":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"xmlrpc":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"xmlwriter":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"xsl":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"yaml":            {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"zbarcode":        {"7.0", "7.1", "7.2", "7.3"},
	"zendopcache":     {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"zip":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
}

var availableServices = []*service{
	{
		Type: "chrome-headless",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"73", "80", "81", "83", "84", "86", "91", "95"},
		},
	},
	{
		Type: "elasticsearch",
		Versions: serviceVersions{
			Deprecated: []string{"0.9", "1.4", "1.7", "2.4", "5.2", "5.4", "6.5", "6.8", "7.2", "7.5", "7.7", "7.9", "7.10"},
			Supported: []string{},
		},
	},
	{
		Type: "influxdb",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.2", "1.3", "1.7", "1.8", "2.2", "2.3"},
		},
	},
	{
		Type: "kafka",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"2.1", "2.2", "2.3", "2.4", "2.5", "2.6", "2.7"},
		},
	},
	{
		Type: "mariadb",
		Versions: serviceVersions{
			Deprecated: []string{"5.5"},
			Supported:  []string{"10.0", "10.1", "10.2", "10.3", "10.4", "10.5", "10.6"},
		},
	},
	{
		Type: "memcached",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.4", "1.5", "1.6"},
		},
	},
	{
		Type: "mongodb",
		Versions: serviceVersions{
			Deprecated: []string{"3.0", "3.2", "3.4", "3.6"},
			Supported: []string{},
		},
	},
	{
		Type: "mongodb-enterprise",
		Versions: serviceVersions{
			Deprecated: []string{"4.0"},
			Supported:  []string{"4.2", "4.4", "5.0"},
		},
	},
	{
		Type: "mysql",
		Versions: serviceVersions{
			Deprecated: []string{"5.5"},
			Supported:  []string{"10.0", "10.1", "10.2", "10.3", "10.4", "10.5", "10.6"},
		},
	},
	{
		Type: "network-storage",
		Versions: serviceVersions{
			Deprecated: []string{"1.0"},
			Supported:  []string{"2.0"},
		},
	},
	{
		Type: "opensearch",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.1", "1.2", "2.0"},
		},
	},
	{
		Type: "oracle-mysql",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"5.7", "8.0"},
		},
	},
	{
		Type: "postgresql",
		Versions: serviceVersions{
			Deprecated: []string{"9.3", "9.4", "9.5"},
			Supported:  []string{"9.6", "10", "11", "12", "13"},
		},
	},
	{
		Type: "rabbitmq",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"3.5", "3.6", "3.7", "3.8", "3.9"},
		},
	},
	{
		Type: "redis",
		Versions: serviceVersions{
			Deprecated: []string{"2.8", "3.0"},
			Supported:  []string{"3.2", "4.0", "5.0", "6.0", "6.2", "7.0"},
		},
	},
	{
		Type: "solr",
		Versions: serviceVersions{
			Deprecated: []string{"3.6", "4.10", "6.3", "6.6", "7.6"},
			Supported:  []string{"7.7", "8.0", "8.4", "8.6", "8.11"},
		},
	},
	{
		Type: "varnish",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"5.1", "5.2", "6.0", "6.3"},
		},
	},
	{
		Type: "vault-kms",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.6", "1.8"},
		},
	},
}
