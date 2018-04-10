var createReactClass = require('create-react-class');

global.Hello = reactCreateClass({
    render: function () {
		var message = this.props.message
        return React.DOM.h1(null, message)
    }

});
