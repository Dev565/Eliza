const form = $("#userInput");
const list = $("#conversation_list");


//adapted from: https://stackoverflow.com/questions/905222/enter-key-press-event-in-javascript
form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }
	
	event.preventDefault(); // stops the refresh
	
	const recieved = form.val();
	form.val(" ");
	
	list.append("<li class='list-group-item text-left list-group-item-success'  id='leftList'>"+ "User: " + recieved + "</li>");
	
	const query = {"userInput" : recieved}
	$.get("/chat", query)
	
		.done(function(resp){
			setTimeout(function(){
				
			
				const newItem = "<li class = 'list-group-item list-group-item-failure'>"+ "Eliza: " +resp+"</li>";
			
				list.append(newItem)}, 1000);
		}).fail(function(){
			const newItem = "<li class = 'list-group-item list-group-item-failure'>The person you are trying to connect to has blocked you.</li>";
			list.append(newItem);
		});
});

/*
// atempt of using ajax
// adapted from: https://github.com/ET-CS/golang-response-examples/blob/master/ajax-json.go
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    //parse request to struct
    var d Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
*/