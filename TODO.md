# Tasks
* I have to finish my minimap: currently it tries to re-render whenever the room changes, but it tries to do it with a hardcoded url to get the current minimap. This won't work as it comes from the backend, so I'm sure I'll need some way of keeping it as application state (maybe with hyperscript?) Also, I don't even think the `changedRoom` event is triggering, so I should check it
* Actually exiting the maze
* Showing whether it's an entrance in the room
* Implement JSON+HAL
* Implement Maze+XML
* It'd be great if I were able to have a base template for pages (not requested with htmx) instead of copy-pasting most of the template like I do for `index`, `room` and `maze`
