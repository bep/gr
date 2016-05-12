require=(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({"ElapserExtModule":[function(require,module,exports){
 module.exports = React.createClass({
    render: function () {
        var message =
            'Module JavaScript timer has been successfully running for ' + this.props.elapsed + ' seconds.';

        return React.DOM.div(null, message)
    },

    shouldComponentUpdate: function (nextProps, nextState) {
        return  nextProps.elapsed !== this.props.elapsed;
    }

});

},{}]},{},[]);
