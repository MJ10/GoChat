new Vue({
  el: '#app',

  data: {
    ws: null,
    newMsg: '',
    chatContents: '',
    email: null,
    username: null,
    joined: false
  },
  created: function() {
    var self = this;
    this.ws = new WebSocket('ws://' + window.location.host + '/ws');
    this.ws.addEventListener('message',  function(e) {
      var msg = JSON.parse(e.data);
      self.chatContents += '<div class="chip">'+ msg.username + '</div>'
                            + emojione.toImage(msg.message);
      document.getElementById('chat-messages').scrollTop = document.getElementById('chat-messages').scrollHeight;

    });
  },
  methods: {
    send: function() {
      if (this.newMsg != null) {
        this.ws.send(JSON.stringify({
          email: this.email,
          username: this.username,
          message: $('<p>').html(this.newMsg).text()
        }))
      }
      this.newMsg = '';
    },
    join: function() {
      if(!this.email) {
        Materialize.toast('You must enter an email', 2000);
        return;
      }
      if (!this.username) {
        Materialize.toast('You must enter a username', 2000);
        return;
      }
      this.email = $('<p>').html(this.email).text();
      this.username = $('<p>').html(this.username).text();
      this.joined = true;
    }
  }
})
