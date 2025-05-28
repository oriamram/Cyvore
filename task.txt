Home Assignment: Full-Stack
Development Task for Cyvore
Overview:
At Cyvore, we focus on improving workspace security. For this task, we would like you to
integrate a cybersecurity reconnaissance tool into a full-stack web application. You will be
able to choose from one of the following three open-source tools for the assignment:
1. Amass - for subdomain enumeration and OSINT gathering.
2. EyeWitness - for website reconnaissance, including screenshots and metadata
extraction.
3. Aquatone - for website screenshot generation and metadata gathering.
Task:
● Implement the backend
Use Gin (GO) to integrate one of the tools.
The backend should run the selected tool, capture the JSON output, and expose it via a
REST API.
● Develop the frontend
Use Vue 3, Vite, TypeScript, Tailwind CSS, and ShadCN to visualize the data collected
by the tool.
The frontend should present the results in an organized and interactive way.
Requirements:
● Backend: Use the selected tool to gather data (e.g., subdomains, websites, open ports)
and expose it through a REST API.
● Frontend: Build a dashboard containing a login page and a table view to display the
data. Allow users to:
○ View results (e.g., list subdomains, websites, or open ports).
○ Filter or search through results.
○ Show detailed information (e.g., metadata, IP addresses, vulnerabilities).

Bonus (Optional):
● Implement authentication (using JWT or API keys).
● Add real-time updates via WebSockets or polling to keep the data fresh.

Deliverables:
1. A GitHub repository with your code.
2. A README file with setup instructions and an overview of your design choices.
3. A brief explanation of the tool you chose, how you integrated it, and the structure of your
application.

Evaluation Criteria:
● Backend: The integration of the selected tool and handling of the JSON output in a
clean, RESTful API.
● Frontend: The clarity, usability, and responsiveness of the UI, along with how well the
data is presented.
● Code Quality: Maintainability, organization, and documentation of the code.
● Bonus Features: Implementing authentication or real-time updates.

Time Estimate:
● We expect this task to take 8-10 hours to complete. Please manage your time and
submit when you&#39;re satisfied with your work.