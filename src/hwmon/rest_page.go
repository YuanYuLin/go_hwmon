package hwmon

import "config"
import "net/http"
import "fmt"

func PageIndex(w http.ResponseWriter, r* http.Request) {
        fmt.Fprintf(w, "<html><title>Index</title><body><a href='/debug.html'>Debug</a></body></html>")
}

func PageTimer(w http.ResponseWriter) {
	fmt.Fprintln(w, "var timer = setInterval(ShowMeDebug, 1000);")
}

func PageGenTable(w http.ResponseWriter) {
	fmt.Fprintln(w, "var tbl = document.createElement('table');")
	fmt.Fprintln(w, "tbl.style.width='100%';")
	fmt.Fprintln(w, "tbl.setAttribute('border', '1');")
	fmt.Fprintln(w, "for(var r=0; r<10; r++) {")
	fmt.Fprintln(w, "	var tr = document.createElement('tr');")
	fmt.Fprintln(w, "	for(var c=0; c<100; c++) {")
	fmt.Fprintln(w, "		if (r === 0) {")
	fmt.Fprintln(w, "			var th = document.createElement('th');")
	fmt.Fprintln(w, "			th.innerHTML = 'inst : ' + c;")
	fmt.Fprintln(w, "			tr.appendChild(th);")
	fmt.Fprintln(w, "		} else {")
	fmt.Fprintln(w, "			var td = document.createElement('td');")
	fmt.Fprintln(w, "			var el = document.createElement('p');")
	fmt.Fprintln(w, "			el.setAttribute('id', 'F_' + r + '_' + c);")
	fmt.Fprintln(w, "			td.appendChild(el);")
	fmt.Fprintln(w, "			tr.appendChild(td);")
	fmt.Fprintln(w, "		};")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	tbl.appendChild(tr);")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "document.body.appendChild(tbl);")
}

func PageRequestSensor(w http.ResponseWriter) {
	fmt.Fprintln(w, "function ShowMeDebug(){")
	var el_name string

	url_abstemp := fmt.Sprintf("/api/v1/hwmon/get/device/abstemp")
	el_name = fmt.Sprintf("%s_%d_%d", "F", config.ENTITY_AMB, 1)
	fmt.Fprintf(w, "	request_sensor_reading('%s', '%s', %d, %d);", el_name, url_abstemp, config.ENTITY_AMB, 1)
	fmt.Fprintln(w, "")
/*
	el_name = fmt.Sprintf("%s_%d_%d", "F", config.ENTITY_CPU, 1)
	fmt.Fprintf(w, "	request_sensor_reading('%s', '%s', %d, %d);", el_name, url_abstemp, config.ENTITY_CPU, 1)
	fmt.Fprintln(w, "")

	el_name = fmt.Sprintf("%s_%d_%d", "F", config.ENTITY_CPU, 2)
	fmt.Fprintf(w, "	request_sensor_reading('%s', '%s', %d, %d);", el_name, url_abstemp, config.ENTITY_CPU, 2)
	fmt.Fprintln(w, "")
*/
	url_reltemp := fmt.Sprintf("/api/v1/hwmon/get/device/reltemp")
	el_name = fmt.Sprintf("%s_%d_%d", "F", config.ENTITY_CPU, 1)
	fmt.Fprintf(w, "	request_sensor_reading('%s', '%s', %d, %d);", el_name, url_reltemp, config.ENTITY_CPU, 1)
	fmt.Fprintln(w, "")

	el_name = fmt.Sprintf("%s_%d_%d", "F", config.ENTITY_CPU, 2)
	fmt.Fprintf(w, "	request_sensor_reading('%s', '%s', %d, %d);", el_name, url_reltemp, config.ENTITY_CPU, 2)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "}")
}

func PageLibs(w http.ResponseWriter) {
	fmt.Fprintln(w, "function request_sensor_reading(el_name, url, eid, idx){")
	fmt.Fprintln(w, "	var json_data = {entity:eid, instant:idx};")
	fmt.Fprintln(w, "	var el = document.getElementById(el_name);")
	fmt.Fprintln(w, "	var xhr = new XMLHttpRequest();")
	fmt.Fprintln(w, "	xhr.open('POST', url);")
	fmt.Fprintln(w, "	xhr.setRequestHeader('content-type', 'application/json');")
	fmt.Fprintln(w, "	xhr.onreadystatechange = function() {")
	fmt.Fprintln(w, "		if (xhr.readyState == 4) {")
	fmt.Fprintln(w, "			if(xhr.getResponseHeader('content-type')==='application/json'){")
	fmt.Fprintln(w, "				var response = JSON.parse(xhr.responseText);")
	fmt.Fprintln(w, "				var dev = response.data;")
	fmt.Fprintln(w, "				if(dev.valuetype === 0xF0) {")
	fmt.Fprintln(w, "					el.innerHTML = 'NA';")
	fmt.Fprintln(w, "				} else {")
	fmt.Fprintln(w, "					el.innerHTML = dev.value;")
	fmt.Fprintln(w, "				}")
	fmt.Fprintln(w, "			}")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "    xhr.send(JSON.stringify(json_data));")
	fmt.Fprintln(w, "}")
}

func PageDebug(w http.ResponseWriter, r* http.Request) {
    fmt.Fprintln(w, "<html><body><script>")
    
	PageLibs(w)
	PageGenTable(w)
	PageTimer(w)
	PageRequestSensor(w)

    fmt.Fprintln(w, "</script></body></html>")
}
