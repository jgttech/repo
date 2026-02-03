# Uses AI to generate a commit message.
[group('claude')]
save:
  echo "Saving changes, please wait.."
  just claude prompt save
  echo "Done"
