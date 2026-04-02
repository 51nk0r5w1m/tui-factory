# Agent Instructions: tui-factory

## Objective
Build a reusable Bubble Tea TUI factory that serves as:
- a terminal UI design system
- a component library
- a pattern engine for operator workflows

The first consumer app will be `tenantui`.

## Architecture Rules

### 1. Separation of Concerns
- `tui-factory` = UI only
- `tenantui` = AWS, Terraform, and business logic
- Never mix them

### 2. Core Layers

#### Primitives
- table
- form
- modal
- status
- progress
- logs

#### Patterns
- list/detail
- wizard
- dashboard
- task runner

#### Apps
- `tenantui` as the first consumer

### 3. State Model
- Every screen is a Bubble Tea model
- No global mutable state
- Explicit transitions only

### 4. Async Rules
- Always show loading
- Always show progress
- Always show failure state

## Initial Tasks
1. Create app shell
2. Implement theme system
3. Implement keymap system
4. Build table component
5. Build form component
6. Build modal system
7. Implement wizard pattern

## Coding Style
- Idiomatic Go
- Small packages
- No over-engineering
- Prefer clarity over abstraction

## Execution Mode
When given a task:
1. Create necessary files
2. Implement a minimal working version
3. Ensure it compiles
4. Add concise comments where needed
5. Do not over-optimize

## Important
This project is not yet:
- an AWS tool
- a Terraform wrapper

It is a UI system first.
