/**
 * This file provided by Facebook is for non-commercial testing and evaluation
 * purposes only. Facebook reserves all rights not expressly granted.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * FACEBOOK BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
function padding(num, length) {
  for(var len = (num + "").length; len < length; len = num.length) {
      num = "0" + num;            
  }
  return num;
}
var Log = React.createClass({
  rawMarkup: function() {
    var md = new Remarkable();
    var logstr = padding(this.props.log.Key, 4) + `:[${this.props.log.Name}] ${this.props.log.Text}`
    var rawMarkup = md.render(logstr);
    return { __html: rawMarkup };
  },
  render: function() {
    return (
        <div className="log">
            <span dangerouslySetInnerHTML={this.rawMarkup()}/>
        </div>
  );
}
});

var LogBox = React.createClass({
  getInitialState: function() {
    return {data: []};
  },
  componentWillMount: function() {
      this.setState({data: []});
  },
  componentDidMount: function() {
    SubscribeToLog((log, err) => {
        if (err) {
            console.log(err);
            return;
        }
        let index = this.state.data.length + 1;
        if (10000 < index) {
            log.Key = 1
            this.setState({data: [log]})
            return
        }
        log.Key = index
        let newdata = this.state.data.concat(log);
        this.setState({data: newdata});
    });
  },
  render: function() {
    return (
      <div className="logbox">
        <h1>Log Monitor</h1>
        <LogList data={this.state.data} />
      </div>
    );
  }
});

var LogList = React.createClass({
  scrollToBottom: function() {
    let el = document.getElementById('scrolldiv')
    el.scrollIntoView({ behavior: "smooth" });
  },
  componentDidMount: function() {
    this.scrollToBottom();
  },
  componentDidUpdate: function() {
      this.scrollToBottom();
  },
  render: function() {
    var loglistFunc = this.props.data.map(function(log) {
      return (
        <Log log={log} key={log.Key}/>
      );
    });
    return (
      <div>
        <div className="loglist" >
          {loglistFunc}
        </div>
        <div id='scrolldiv'/>
      </div>
    );
  }
});

ReactDOM.render(
  <LogBox pollInterval={2000} />,
  document.getElementById('content')
);