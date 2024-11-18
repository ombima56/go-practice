# Instrunction
To-Do List Manager with Priority Sorting
**Prompt**
Design a To-Do List Manager in Go that allows users to add, view, delete, and sort tasks by priority. The application should be interactive and utilize structs, slices, and functions for modular design.
## Requirements
## 1. Task Struct:
   * Define a struct, Task, with the following fields:
      * ID: A unique integer identifier for the task.
      * Description: A string containing the task details.
      * Priority: An integer (1 = High, 2 = Medium, 3 = Low).
## 2. Core Features:
   * Add Task: Add a new task with a description and priority.
   * View Tasks: Display all tasks sorted by priority (High to Low). Include the task ID, description, and priority level.
   * Delete Task: Delete a task by its unique ID.
   * Update Task: Update the description or priority of an existing task by its ID.
## 3. Interactive CLI:
   * Create a simple command-line interface with the following options:
      * 1 to Add Task.
      * 2 to View Tasks.
      * 3 to Delete Task.
      * 4 to Update Task.
      * 5 to Exit the program.
## 4. Sorting:
   * Implement a function to sort tasks by priority. High-priority tasks should appear first.
## 5. Validation:
   * Ensure task descriptions are not empty.
   * Validate that priority values are within the accepted range (1-3).
   * Handle invalid user inputs gracefully.
## 6. Storage:
   * Use an in-memory slice to store tasks. Tasks should be appended when added and dynamically removed when deleted.
## Testing
1. Add at least three tasks with varying priorities.
2. View tasks to confirm they are sorted correctly by priority.
3. Delete a task and view the updated list to ensure it was removed.
4. Update a task’s description and/or priority and verify the changes.
