# AIServer

AIServer is a chat server written in Go using the Gin framework. It allows you to query an Ollama server with a JSON object containing participants and interview data.

## Prerequisites

- Go (version 1.23 or higher)
- Ollama server installed and available in your PATH

## Setup

Clone the repository:

   ```
   git clone https://github.com/akmal/aiserver.git
   cd aiserver
   ```

## Usage
### Running the Server
Start the server with the desired LLM model (default is llama3.1):

```
    ./aiserver -model=mistral-nemo
```

You can also enable debug mode with the -debug flag:

```
    ./aiserver -model=llama3.1 -debug
```
Disregard the '_Ollama exited with error: exit status 1_' message

### Running the Client
Create an input JSON file with the desired participants and interview data. For example, create a file named input.json with the following content:

```
{
  "participants": [
    {
      "name": "Ethan Thompson",
      "age": 29,
      "gender": "male",
      "occupation": "Graphic Designer",
      "location": "Los Angeles, CA",
      "marital_status": "unmarried"
    },
    {
      "name": "Sophia Rodriguez",
      "age": 27,
      "gender": "female",
      "occupation": "Event Coordinator",
      "location": "Miami, FL",
      "marital_status": "unmarried"
    }
  ],
  "interview": [
    {
      "question": "What does intimacy mean to you, both physically and emotionally?",
      "Ethan Thompson": "For me, intimacy is about feeling connected with someone on a deep level. It's not just physical; it's emotional too. When I'm with someone, I want to feel like we can be vulnerable together,and share our true selves.",
      "Sophia Rodriguez": "Intimacy means being able to trust someone completely and feeling comfortable in their presence. It's about creating a safe space where we can both be ourselves without fear of judgment."
    },
    {
      "question": "How do you handle change or unexpected events?",
      "Ethan Thompson": "I try to stay calm and adapt quickly. I think about what needs to happen next and prioritize tasks. Sometimes, taking a step back and re-evaluating the situation helps me find a solution.",
      "Sophia Rodriguez": "I'm a bit of a planner, so when something unexpected happens, it can throw me off. But I try to take a deep breath, assess the situation, and then adjust my plans accordingly. Communication with others is key in these situations.  "
    },
    {
      "question": "What does trust mean to you in a relationship?",
      "Ethan Thompson": "Trust means being able to rely on someone without question. It's knowing that they'll be there for me no matter what, and that we can communicate openly about everything.",
      "Sophia Rodriguez": "To me, trust is about feeling secure and supported in the relationship. It means being able to share my thoughts, feelings, and desires with my partner without fear of judgment or rejection.  "
    },
    {
      "question": "How do you think we can best navigate disagreements and compromises in our relationship?",
      "Ethan Thompson": "We should communicate openly and honestly about what we want and need. Sometimes, it's not about getting what I want; it's about finding a solution that works for both of us.",
      "Sophia Rodriguez": "I think active listening is key. We need to really hear each other out and try to understand the other person's perspective before responding or reacting."
    },
    {
      "question": "What are your expectations from your partner in terms of emotional support?",
      "Ethan Thompson": "I expect my partner to be understanding and supportive when I'm going through a tough time. I want them to listen to me without judgment and offer guidance when needed.",
      "Sophia Rodriguez": "I expect my partner to be empathetic and validating. I want them to acknowledge my feelings and concerns, and offer support and encouragement when I need it."
    },
    {
      "question": "How do you prioritize self-care and personal growth in your life?",
      "Ethan Thompson": "I make time for activities that nourish my mind, body, and soul. This includes exercise, meditation, and spending time with loved ones.",
      "Sophia Rodriguez": "I prioritize self-care by setting aside time for myself each day. This can be as simple as taking a relaxing bath or reading a book. I also make sure to schedule activities that bring me joy and help me grow as a person.  "
    },
    {
      "question": "What are your thoughts on conflict resolution in relationships?",
      "Ethan Thompson": "I think conflicts are inevitable, but it's how we resolve them that matters. We should communicate openly and honestly, and be willing to listen to each other's perspectives.",
      "Sophia Rodriguez": "I believe conflicts can actually strengthen a relationship if handled correctly. It's about being able to navigate difficult conversations with empathy and understanding."
    }
  ]
}

```
Run the client with the input JSON file:

```
	./client input.json

```

You can also enable debug mode by appending '?debug=true' to the input JSON file:

```
	./client input.json?debug=true
```
