{{define "body"}}

  <script src="/js/todo.js" type="riot/tag"></script>

  <!-- mount the tag -->
  <script>
   riot.mount('todo'
	      , {
      title: 'I want to behave!',
      items: [
        { title: 'Avoid excessive caffeine', done: true },
        { title: 'Hidden item',  hidden: true },
        { title: 'Be nice to people' }
      ]}

   );

  </script>

  <todo></todo>
{{end}}
