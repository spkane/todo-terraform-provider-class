output "todo_1_ids" {
  value = "${todo.test1.*.id}"
}

output "imported_id" {
  value = "${todo.imported.*.id}"
}

output "todo_2_ids" {
  value = "${todo.test2.*.id}"
}
