// some basic for resource leaks:
//
// graphs:
// connection amount: {{project}}_http_connections
// too many qps: rate({{project}}_http_handler_ttl_count[1m])*60
// long connection: histogram_quantile(0.999, sum(rate({{project}}_http_handler_ttl_bucket[5m])) by (le))
// db connection: {{project}}_{{db_name}}_db_connection{type="open"}
// inuse db connection: {{project}}_{{db_name}}_db_connection{type="inuse"}
// cpu: sum by﻿ ﻿(﻿instance﻿)﻿({{project}}_system_cpu﻿)
// system mem: {{project}}_system_mem{cat="virtual_total"}
// process mem: {{project}}_process_mem{cat="sys"}
// heap mem: {{project}}_process_mem{cat="heap_inuse"}
// load average 1: {{project}}_system_load{﻿cat﻿=﻿"avg_1"﻿}
// load average 5: {{project}}_system_load{﻿cat﻿=﻿"avg_5"﻿}
// load average 15: {{project}}_system_load{﻿cat﻿=﻿"avg_15"﻿}
//
// calculated:
// long connection
// connection amount
// db inuse
// load average 1 5 15 value 0.7 good 1 ok 5 horrible