
var start = new Date().getTime()

setInterval(function () {
    ReactDOM.render(
	    Elapser({
            elapsed: Math.round((( new Date().getTime() - start ) / 1000)),
			title: "Interop from JavaScript"
        }),
        document.getElementById('react-reverse')
    );
}, 50);
