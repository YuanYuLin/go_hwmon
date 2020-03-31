package hwmon

import "config"
import "net/http"
import "fmt"

var MAX_ENTITY int32 = 50
var MAX_INSTANT int32 = 100
func PageIndex(w http.ResponseWriter, r* http.Request) {
        fmt.Fprintf(w, "<html><title>Index</title><body>")
	fmt.Fprintf(w, "<a href='/debug.html'>Debug</a><br/>")
	fmt.Fprintf(w, "<a href='/devicefanmap.html'>DeviceFanMap</a><br/>")
	fmt.Fprintf(w, "</body></html>")
}

func PageTimer(w http.ResponseWriter, js_func_name string, microseconds int32) {
	fmt.Fprintf (w, "var timer_%s = setInterval(%s, %d);\n", js_func_name, js_func_name, microseconds)
}

func PageGenTable(w http.ResponseWriter) {
	fmt.Fprintln(w, "function lookup_entity(eid){")
	fmt.Fprintf (w, "	var e001 = %d; if (eid == e001) return e001 + ':CPU';", config.ENTITY_PROCESSOR)
	fmt.Fprintf (w, "	var e002 = %d; if (eid == e002) return e002 + ':AMB';", config.ENTITY_EXTERNAL_ENVIROMENT)
	fmt.Fprintf (w, "	var e003 = %d; if (eid == e003) return e003 + ':AIC';", config.ENTITY_ADD_IN_CARD)
	fmt.Fprintf (w, "	var e004 = %d; if (eid == e004) return e004 + ':FAN';", config.ENTITY_FAN_COOLING)
	fmt.Fprintf (w, "	var e005 = %d; if (eid == e005) return e005 + ':DIMM';", config.ENTITY_MEMORY_DEVICE)
	fmt.Fprintln(w, "	return 'NA';")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "function genInstantTable(entity, instant){")
	fmt.Fprintln(w, "	var el_t = document.createElement('p');")
	fmt.Fprintln(w, "	el_t.setAttribute('id', 'T_' + entity + '_' + instant);")
	fmt.Fprintln(w, "	if (instant === 0) {")
	fmt.Fprintln(w, "		el_t.innerHTML=lookup_entity(entity);")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	var el_f = document.createElement('p');")
	fmt.Fprintln(w, "	el_f.setAttribute('id', 'F_' + entity + '_' + instant);")

	fmt.Fprintln(w, "	var tbl = document.createElement('table');")
	fmt.Fprintln(w, "	tbl.setAttribute('border', '0');")
	fmt.Fprintln(w, "	var tr = document.createElement('tr');")
	fmt.Fprintln(w, "	var tr2 = document.createElement('tr');")
	fmt.Fprintln(w, "	var td = document.createElement('td');")
	fmt.Fprintln(w, "	var td2 = document.createElement('td');")
	fmt.Fprintln(w, "	td.appendChild(el_t);")
	fmt.Fprintln(w, "	td2.appendChild(el_f);")
	fmt.Fprintln(w, "	tr.appendChild(td);")
	fmt.Fprintln(w, "	tr2.appendChild(td2);")
	fmt.Fprintln(w, "	tbl.appendChild(tr);")
	fmt.Fprintln(w, "	tbl.appendChild(tr2);")
	fmt.Fprintln(w, "	return tbl;")
	fmt.Fprintln(w, "}")

	fmt.Fprintln(w, "var tbl = document.createElement('table');")
	fmt.Fprintln(w, "tbl.style.width='100%';")
	fmt.Fprintln(w, "tbl.setAttribute('border', '1');")
	fmt.Fprintf (w, "var MAX_E=%d, MAX_I=%d;\n", MAX_ENTITY, MAX_INSTANT)
	fmt.Fprintln(w, "for(var r=0; r<MAX_E; r++) {")
	fmt.Fprintln(w, "	var tr = document.createElement('tr');")
	fmt.Fprintln(w, "	for(var c=0; c<MAX_I; c++) {")
	fmt.Fprintln(w, "		if (r === 0) {")
	fmt.Fprintln(w, "			var th = document.createElement('th');")
	fmt.Fprintln(w, "			th.innerHTML = c;")
	fmt.Fprintln(w, "			tr.appendChild(th);")
	fmt.Fprintln(w, "		} else {")
	fmt.Fprintln(w, "			var td = document.createElement('td');")
	fmt.Fprintln(w, "			td.appendChild(genInstantTable(r, c));")
	fmt.Fprintln(w, "			tr.appendChild(td);")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	tbl.appendChild(tr);")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "var el = document.getElementById('main_ctx');")
	fmt.Fprintln(w, "el.appendChild(tbl);")
}

func PageRequestSensors(w http.ResponseWriter) (string){
	func_name := "RequestSensors"
	url_abstemp := "/api/v1/hwmon/get/device/abstemp"
	url_reltemp := "/api/v1/hwmon/get/device/reltemp"
	url_expectduty := "/api/v1/hwmon/get/map/allexpectduty"
	url_dutyout := "/api/v1/hwmon/get/map/allfandutyout"

	fmt.Fprintln (w, "var g_req_list = [")
	fmt.Fprintf (w, "{'en':1, 'url':'%s', 'eid':%d, 'inst':%d, 'callback':%s},",
				url_abstemp, config.ENTITY_EXTERNAL_ENVIROMENT, -1, "pushListToTableTemp")
	fmt.Fprintf (w, "{'en':1, 'url':'%s', 'eid':%d, 'inst':%d, 'callback':%s},",
				url_abstemp, config.ENTITY_ADD_IN_CARD, -1, "pushListToTableTemp")
	fmt.Fprintf (w, "{'en':1, 'url':'%s', 'eid':%d, 'inst':%d, 'callback':%s},",
				url_abstemp, config.ENTITY_MEMORY_DEVICE, -1, "pushListToTableTemp")
	fmt.Fprintf (w, "{'en':1, 'url':'%s', 'eid':%d, 'inst':%d, 'callback':%s},",
				url_reltemp, config.ENTITY_PROCESSOR, -1, "pushListToTableTemp")
	fmt.Fprintf (w, "{'en':1, 'url':'%s', 'eid':%d, 'inst':%d, 'callback':%s},",
				url_expectduty, -1, -1, "pushListToTableExpectDuty")
	fmt.Fprintf (w, "{'en':1, 'url':'%s', 'eid':%d, 'inst':%d, 'callback':%s},",
				url_dutyout, -1, -1, "pushListToTableDutyOut")
	fmt.Fprintln(w, "{'en':0, 'url':'', 'eid':-1, 'inst':-1, 'callback':''}]")
	fmt.Fprintln(w, "var g_req_idx = 0;")
	fmt.Fprintf (w, "function %s(){", func_name)
	fmt.Fprintln(w, "	if(g_req_idx >= g_req_list.length) {")
	fmt.Fprintln(w, "		g_req_idx = 0;")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	var obj = g_req_list[g_req_idx];")
	fmt.Fprintln(w, "	if(obj.en)")
	fmt.Fprintln(w, "		request_to(obj.url, obj.eid, obj.inst, obj.callback);")
	fmt.Fprintln(w, "	g_req_idx+=1;")
	fmt.Fprintln(w, "}")
	return func_name
}

func PageRequestDeviceFanMap(w http.ResponseWriter) (string) {
	func_name := "RequestDeviceFanMap"
	url_devicefanmap := "/api/v1/hwmon/get/device/fanmap"
	fmt.Fprintln(w, "var g_idx_E=1;")
	fmt.Fprintln(w, "var g_idx_I=1;")
	fmt.Fprintf (w, "function %s(){", func_name)
	fmt.Fprintln(w, "")
	fmt.Fprintf (w, "	if(g_idx_I >= %d){", MAX_INSTANT)
	fmt.Fprintln(w, "		g_idx_E +=1;")
	fmt.Fprintln(w, "		g_idx_I = 1;")
	fmt.Fprintln(w, "	}")
	fmt.Fprintf (w, "	if(g_idx_E >= %d){", MAX_ENTITY)
	fmt.Fprintln(w, "		g_idx_E = 1;")
	fmt.Fprintln(w, "	}")
	fmt.Fprintf (w, "	request_to('%s', g_idx_E, g_idx_I, pushListToTableFan);", url_devicefanmap)
	fmt.Fprintln(w, "	g_idx_I += 1;")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")

	return func_name
}

func PageLibsToTable(w http.ResponseWriter) {
	fmt.Fprintln(w, "function pushListToTableFan(list) {")
	fmt.Fprintln(w, "	var text = '';")
	fmt.Fprintln(w, "	var el_name = '';")
	fmt.Fprintln(w, "	var found = 0;")
	fmt.Fprintln(w, "	for(var key in list){")
	fmt.Fprintln(w, "		var dev = list[key];")
	fmt.Fprintln(w, "		if(dev.valuetype === 0xF0) {")
	fmt.Fprintln(w, "			//el.innerHTML = 'NA';")
	fmt.Fprintln(w, "		} else {")
	fmt.Fprintln(w, "			el_name = 'F_' + dev.entity + '_' + dev.instant;")
	fmt.Fprintln(w, "			text +=dev.value +',';")
	fmt.Fprintln(w, "			found = 1;")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	if(found) {")
	fmt.Fprintln(w, "		var el = document.getElementById(el_name);")
	fmt.Fprintln(w, "		el.innerHTML = text;")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "function pushListToTableDutyOut(list) {")
	fmt.Fprintln(w, "	for(var key in list){")
	fmt.Fprintln(w, "		var dev = list[key];")
	fmt.Fprintln(w, "		if(dev.valuetype === 0xF0) {")
	fmt.Fprintln(w, "			//el.innerHTML = 'NA';")
	fmt.Fprintln(w, "		} else {")
	fmt.Fprintln(w, "			el_name = 'F_' + dev.entity + '_' + dev.instant;")
	fmt.Fprintln(w, "			var el = document.getElementById(el_name);")
	fmt.Fprintln(w, "			el.innerHTML = dev.value;")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "function pushListToTableExpectDuty(list) {")
	fmt.Fprintln(w, "	var debug_ctx = document.getElementById('debug_ctx1');")
	fmt.Fprintln(w, "	debug_ctx.innerHTML = '';")
	fmt.Fprintln(w, "	for(var key in list){")
	fmt.Fprintln(w, "		var dev = list[key];")
	fmt.Fprintln(w, "		if(dev.valuetype === 0xF0) {")
	fmt.Fprintln(w, "			//el.innerHTML = 'NA';")
	fmt.Fprintln(w, "		} else {")
	fmt.Fprintln(w, "			debug_ctx.innerHTML += '[Entity-Instant-Duty][' + dev.entity + '-' + dev.instant + '-' + dev.value + ']';")
	fmt.Fprintln(w, "			debug_ctx.innerHTML += '[' + dev.key + ']';")
	fmt.Fprintln(w, "			debug_ctx.innerHTML += '<br/>';")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "function pushListToTableTemp(list) {")
	fmt.Fprintln(w, "	for(var key in list){")
	fmt.Fprintln(w, "		var dev = list[key]")
	fmt.Fprintln(w, "		if(dev.valuetype === 0xF0) {")
	fmt.Fprintln(w, "			//el.innerHTML = 'NA';")
	fmt.Fprintln(w, "		} else {")
	fmt.Fprintln(w, "			var el_name = 'T_' + dev.entity + '_' + dev.instant;")
	fmt.Fprintln(w, "			var el = document.getElementById(el_name);")
	fmt.Fprintln(w, "			el.innerHTML = dev.value;")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
}
func PageLibsCommon(w http.ResponseWriter) {
	fmt.Fprintln(w, "function request_to(url, eid, idx, callback){")
	fmt.Fprintln(w, "	var json_data = {entity:eid, instant:idx};")
	fmt.Fprintln(w, "	var xhr = new XMLHttpRequest();")
	fmt.Fprintln(w, "	xhr.open('POST', url);")
	fmt.Fprintln(w, "	xhr.setRequestHeader('content-type', 'application/json');")
	fmt.Fprintln(w, "	xhr.onreadystatechange = function() {")
	fmt.Fprintln(w, "		if (xhr.readyState == 4) {")
	fmt.Fprintln(w, "			if(xhr.getResponseHeader('content-type')==='application/json'){")
	fmt.Fprintln(w, "				var response = JSON.parse(xhr.responseText);")
	fmt.Fprintln(w, "				if(callback) callback(response.data);")
	fmt.Fprintln(w, "			}")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "    xhr.send(JSON.stringify(json_data));")
	fmt.Fprintln(w, "}")
}

func PageDebug(w http.ResponseWriter, r* http.Request) {
    fmt.Fprintln(w, "<html><title>Temperature</title><body>")
    fmt.Fprintln(w, "<div id='main_ctx' style='width:100%;height:70%;overflow:auto'></div>")
    fmt.Fprintln(w, "<div id='debug_ctx1' style='width:47%;height:30%;overflow:auto'></div>")
    fmt.Fprintln(w, "<div id='debug_ctx2' style='width:47%;height:30%;overflow:auto'></div>")
    fmt.Fprintln(w, "<script>")

	PageLibsCommon(w)
	PageLibsToTable(w)
	PageGenTable(w)
	func_name := PageRequestSensors(w)
	PageTimer(w, func_name, 100)

    fmt.Fprintln(w, "</script>")
    fmt.Fprintln(w, "<table><tr><td>Temperature</td></tr><tr><td>Device Map to Fan</td></tr></table>")
    fmt.Fprintln(w, "</body></html>")
}

func PageDeviceFanMap(w http.ResponseWriter, r* http.Request) {
    fmt.Fprintln(w, "<html><title>DeviceFanMap</title><body>")
    fmt.Fprintln(w, "<div id='main_ctx' style='width:100%;height:70%;overflow:auto'></div>")
    fmt.Fprintln(w, "<div id='debug_ctx1' style='width:47%;height:30%;overflow:auto'></div>")
    fmt.Fprintln(w, "<div id='debug_ctx2' style='width:47%;height:30%;overflow:auto'></div>")
    fmt.Fprintln(w, "<script>")

	PageLibsCommon(w)
	PageLibsToTable(w)
	PageGenTable(w)
	func_name := PageRequestDeviceFanMap(w)
	PageTimer(w, func_name, 100)

    fmt.Fprintln(w, "</script>")
    fmt.Fprintln(w, "</body></html>")
}
