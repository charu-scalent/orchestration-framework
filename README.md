# Fault-Tolerant Workflow Orchestration Framework
In a sophisticated distributed system, maintaining consistency across multiple services during complex workflows is paramount. Consider the scenario of an e-commerce system with interconnected services:

 - **Wallet service**: manages user wallet transactions. 
 - **Order service**: handles order processing and tracking. 
 - **Inventory service**: manages product inventory. 
 - **Notification service**: sends notifications to users.

When a user makes a purchase, the Order Service initiates a workflow involving these services. However, failures at any step could compromise system consistency. For instance, a successful wallet deduction followed by a failed inventory update could lead to inconsistency upon retry.

Your task is to design and build a Fault-Tolerant Workflow Orchestration Framework that ensures consistency even when individual steps might fail. This framework should allow retries from specific steps, preventing the entire process from restarting and introducing inconsistencies.

## Considerations:

 - Generate and utilize unique idempotent keys.
 - Implement a mechanism to retry transactions from a specific step.
 - Design idempotent workflow steps.
 - Ensure the framework's flexibility to handle various workflows.
 - Include unit tests to validate the framework's reliability.

This project aims to assess your ability to create a resilient system for orchestrating distributed workflows. Please provide a detailed design of the framework, including the actual implementation and unit tests.
