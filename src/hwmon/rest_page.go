package hwmon

import "net/http"
import "fmt"

func PageIndex(w http.ResponseWriter, r* http.Request) {
        fmt.Fprintf(w, "<html><title>Index</title><body><a href='/debug.html'>Debug</a></body></html>")
}

func PageBegin(w http.ResponseWriter) {
	fmt.Fprintln(w, "<html><body>")
}

func PageEnd(w http.ResponseWriter) {
	fmt.Fprintf(w, "</body></html>")
}

func PageScriptBegin(w http.ResponseWriter) {
	fmt.Fprintln(w, "<script>")
}

func PageScriptEnd(w http.ResponseWriter) {
	fmt.Fprintf(w, "</script>")
}

func PageRequest(w http.ResponseWriter) {
	fmt.Fprintln(w, "{")
	fmt.Fprintln(w, "	var json_data = {entity:2, instant:1};")
	fmt.Fprintln(w, "	var consolebox = document.getElementById('consolebox');")
	fmt.Fprintln(w, "	var xhr = new XMLHttpRequest();")
	fmt.Fprintln(w, "	xhr.open('POST', '/api/v1/hwmon/get/device/abstemp');")
	fmt.Fprintln(w, "	xhr.setRequestHeader('content-type', 'application/json');")
	fmt.Fprintln(w, "	xhr.onreadystatechange = function() {")
	fmt.Fprintln(w, "		if (xhr.readyState == 4) {")
	fmt.Fprintln(w, "			if(xhr.getResponseHeader('content-type')==='application/json'){")
	fmt.Fprintln(w, "				var response = JSON.parse(xhr.responseText);")
	fmt.Fprintln(w, "				var dev = response.data;")
	fmt.Fprintln(w, "				if(dev.valuetype === 0xF0) {")
	fmt.Fprintln(w, "					consolebox.innerHTML += 'Data not found';")
	fmt.Fprintln(w, "				} else {")
	fmt.Fprintln(w, "					consolebox.innerHTML = dev.value;")
	fmt.Fprintln(w, "				}")
	fmt.Fprintln(w, "			} else {")
	fmt.Fprintln(w, "				consolebox.innerHTML += 'Not Json';")
	fmt.Fprintln(w, "			}")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "    xhr.send(JSON.stringify(json_data));")
	fmt.Fprintln(w, "}")
}

func PageDebug(w http.ResponseWriter, r* http.Request) {
        PageBegin(w)
        PageScriptBegin(w)
        
		fmt.Fprintln(w, "  function ShowMe(){")
		PageRequest(w)
		fmt.Fprintln(w, "  }")

		PageScriptEnd(w)
		fmt.Fprintf(w, "<input id='inputbox' type='button' value='click' onclick='ShowMe()'><br>")
		fmt.Fprintf(w, "<p id='consolebox'>_</p>")
        PageEnd(w)
}
