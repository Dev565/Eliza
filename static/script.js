const form = $("#user-input");
const list = $("#conversation_list");



//adapted from: https://stackoverflow.com/questions/905222/enter-key-press-event-in-javascript
form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }
	
	event.preventDefault(); // stops the refresh
	
	const recieved = form.val();
	form.val(" ");
	
	list.append("<li class='list-group-item text-left list-group-item-success'  id='leftList'>" + recieved + "</li>");
	
	const query = {"user-input" : recieved}
	$.get("/chat", query)
	
		.done(function(resp){
			setTimeout(function(){
				list.append(newItem)
			}, 1000);
			
		}).fail(function(){
			const newItem = "<li class = 'list-group-item list-group-item-failure'>THe person you are trying to connect to has blocked you.</li>";
			list.append(newItem);
		});
});