package main
import "fmt"
func main() {
	buffer := make([]byte, 30000)
	ptr := 0
	buffer[ptr] = buffer[ptr] + byte(13)
	for buffer[ptr] != 0 {
		buffer[ptr] = buffer[ptr] - byte(1)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(2)
		ptr = ptr + 3
		buffer[ptr] = buffer[ptr] + byte(5)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(2)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 6
	}
	ptr = ptr + 5
	buffer[ptr] = buffer[ptr] + byte(6)
	ptr = ptr + 1
	buffer[ptr] = buffer[ptr] - byte(3)
	ptr = ptr + 10
	buffer[ptr] = buffer[ptr] + byte(15)
	for buffer[ptr] != 0 {
		for buffer[ptr] != 0 {
			ptr = ptr + 9
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		for buffer[ptr] != 0 {
			ptr = ptr - 9
		}
		ptr = ptr + 9
		buffer[ptr] = buffer[ptr] - byte(1)
	}
	buffer[ptr] = buffer[ptr] + byte(1)
	for buffer[ptr] != 0 {
		ptr = ptr + 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		ptr = ptr + 1
	}
	ptr = ptr - 9
	for buffer[ptr] != 0 {
		ptr = ptr - 9
	}
	ptr = ptr + 8
	for buffer[ptr] != 0 {
		buffer[ptr] = buffer[ptr] - byte(1)
	}
	buffer[ptr] = buffer[ptr] + byte(1)
	ptr = ptr - 7
	buffer[ptr] = buffer[ptr] + byte(5)
	for buffer[ptr] != 0 {
		buffer[ptr] = buffer[ptr] - byte(1)
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 9
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 9
		}
		ptr = ptr + 9
	}
	ptr = ptr + 7
	buffer[ptr] = buffer[ptr] + byte(1)
	ptr = ptr + 27
	buffer[ptr] = buffer[ptr] + byte(1)
	ptr = ptr - 17
	for buffer[ptr] != 0 {
		ptr = ptr - 9
	}
	ptr = ptr + 3
	for buffer[ptr] != 0 {
		buffer[ptr] = buffer[ptr] - byte(1)
	}
	buffer[ptr] = buffer[ptr] + byte(1)
	for buffer[ptr] != 0 {
		ptr = ptr + 6
		for buffer[ptr] != 0 {
			ptr = ptr + 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 2
		}
		ptr = ptr - 9
		for buffer[ptr] != 0 {
			ptr = ptr - 9
		}
		ptr = ptr + 7
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 6
		buffer[ptr] = buffer[ptr] + byte(4)
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 9
		}
		ptr = ptr + 6
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 6
		buffer[ptr] = buffer[ptr] + byte(7)
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 9
		}
		ptr = ptr + 6
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 16
		for buffer[ptr] != 0 {
			ptr = ptr - 9
		}
		ptr = ptr + 3
		for buffer[ptr] != 0 {
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 6
			for buffer[ptr] != 0 {
				ptr = ptr + 7
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 6
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 6
				}
				ptr = ptr - 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 6
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
				}
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 8
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 7
				}
				ptr = ptr - 7
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
				}
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 7
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 7
			}
			ptr = ptr - 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 7
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 5
			}
			ptr = ptr + 9
			buffer[ptr] = buffer[ptr] + byte(15)
			for buffer[ptr] != 0 {
				for buffer[ptr] != 0 {
					ptr = ptr + 9
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
				}
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
						}
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 4
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 9
					}
					ptr = ptr - 8
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 9
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 10
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
					}
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 4
				}
				ptr = ptr - 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 5
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 5
				}
				ptr = ptr - 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 4
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 9
					}
					ptr = ptr - 8
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 9
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 11
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
					}
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 4
				}
				ptr = ptr - 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 36
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 36
				}
				ptr = ptr + 5
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			buffer[ptr] = buffer[ptr] + byte(15)
			for buffer[ptr] != 0 {
				for buffer[ptr] != 0 {
					ptr = ptr + 9
				}
				ptr = ptr - 9
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 21
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 3
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 3
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 13
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
						ptr = ptr + 4
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
					}
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 12
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 8
				}
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr - 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 3
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 4
			}
			ptr = ptr + 9
			buffer[ptr] = buffer[ptr] + byte(26)
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 4
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 4
			}
			ptr = ptr - 4
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 2
			}
			ptr = ptr + 2
			for buffer[ptr] != 0 {
				ptr = ptr - 7
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
					}
					ptr = ptr + 3
				}
				ptr = ptr + 13
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 5
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 6
				for buffer[ptr] != 0 {
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 9
					}
					ptr = ptr + 7
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(15)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 9
						}
						ptr = ptr - 8
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
					}
					ptr = ptr + 9
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 9
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 9
						}
						ptr = ptr - 10
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 7
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 3
					}
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 4
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 9
						}
						ptr = ptr - 8
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
					}
					ptr = ptr + 9
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 9
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 9
						}
						ptr = ptr - 10
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 6
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 6
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 4
					}
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 36
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 36
					}
					ptr = ptr + 5
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 36
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 36
					}
					ptr = ptr + 6
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(15)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 8
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 7
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 7
					}
					ptr = ptr - 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 7
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 3
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 5
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(2)
						ptr = ptr - 4
					}
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
				}
				ptr = ptr - 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
				}
				ptr = ptr - 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 5
				}
				ptr = ptr + 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 6
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 3
							}
							ptr = ptr - 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 12
								for buffer[ptr] != 0 {
									ptr = ptr - 9
								}
								ptr = ptr + 3
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 6
								for buffer[ptr] != 0 {
									ptr = ptr + 9
								}
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 1
							}
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 11
								for buffer[ptr] != 0 {
									ptr = ptr - 9
								}
								ptr = ptr + 4
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 5
								for buffer[ptr] != 0 {
									ptr = ptr + 9
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 1
							}
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr - 8
						}
						ptr = ptr + 8
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 2
							}
							ptr = ptr + 8
						}
						ptr = ptr - 8
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 5
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 4
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 4
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 14
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 11
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr + 3
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr - 3
									}
									ptr = ptr - 1
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 14
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 11
								}
								ptr = ptr - 2
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 4
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 14
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 11
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
							}
							ptr = ptr - 12
						}
						ptr = ptr + 4
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						ptr = ptr - 4
					}
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
					}
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 1
							}
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 1
							}
							ptr = ptr + 8
						}
						ptr = ptr - 8
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 5
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 14
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 10
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr + 4
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr - 4
									}
									ptr = ptr + 1
								}
								ptr = ptr - 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 4
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 14
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 10
								}
								ptr = ptr - 1
							}
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 4
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 4
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 14
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 10
								}
								ptr = ptr + 1
							}
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 4
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 4
							}
							ptr = ptr - 11
						}
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
					}
				}
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
				}
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 5
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 4
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 4
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 14
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr - 3
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 14
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
							}
							ptr = ptr - 2
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 4
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 14
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
							}
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 3
						}
						ptr = ptr - 12
					}
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 6
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(15)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 5
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 9
						}
						ptr = ptr - 8
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
					}
					ptr = ptr + 9
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 9
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 9
						}
						ptr = ptr - 10
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 7
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 3
					}
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 36
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 36
					}
					ptr = ptr + 6
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] + byte(15)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 4
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 4
						}
						ptr = ptr - 4
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 4
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 13
							for buffer[ptr] != 0 {
								ptr = ptr - 9
							}
							ptr = ptr + 4
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 5
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 12
							for buffer[ptr] != 0 {
								ptr = ptr - 9
							}
							ptr = ptr + 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 6
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr - 8
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 3
				}
				ptr = ptr - 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 3
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 6
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 3
						}
						ptr = ptr + 8
					}
					ptr = ptr - 8
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 10
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 12
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 2
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 2
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 10
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 12
							}
							ptr = ptr - 3
						}
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 10
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 12
							}
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
						}
						ptr = ptr - 13
					}
				}
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
				}
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
						}
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 2
						}
						ptr = ptr + 8
					}
					ptr = ptr - 8
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 10
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 1
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 1
								}
								ptr = ptr + 1
							}
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 10
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
							}
							ptr = ptr - 2
						}
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 10
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
							}
							ptr = ptr + 1
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
						}
						ptr = ptr - 12
					}
					ptr = ptr + 5
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 5
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 4
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					ptr = ptr + 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
					}
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 5
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] + byte(2)
						ptr = ptr - 4
					}
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
				}
				ptr = ptr - 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
				}
				ptr = ptr - 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 5
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 11
								for buffer[ptr] != 0 {
									ptr = ptr - 9
								}
								ptr = ptr + 4
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 5
								for buffer[ptr] != 0 {
									ptr = ptr + 9
								}
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 1
							}
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 3
							}
							ptr = ptr - 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 12
								for buffer[ptr] != 0 {
									ptr = ptr - 9
								}
								ptr = ptr + 3
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 6
								for buffer[ptr] != 0 {
									ptr = ptr + 9
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
								}
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 1
							}
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr - 8
						}
						ptr = ptr + 8
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
					}
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 1
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 1
							}
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 1
							}
							ptr = ptr + 8
						}
						ptr = ptr - 8
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 4
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 2
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 2
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 13
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 10
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr + 3
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr - 3
									}
									ptr = ptr + 1
								}
								ptr = ptr - 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 13
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 10
								}
								ptr = ptr - 1
							}
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 13
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 10
								}
								ptr = ptr + 1
							}
							ptr = ptr - 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
							}
							ptr = ptr - 11
						}
						ptr = ptr + 5
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 7
						}
						ptr = ptr - 7
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 5
						}
					}
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 2
							}
							ptr = ptr + 8
						}
						ptr = ptr - 8
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 4
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 3
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 3
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 13
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 11
									for buffer[ptr] != 0 {
										buffer[ptr] = buffer[ptr] - byte(1)
										ptr = ptr + 2
										buffer[ptr] = buffer[ptr] + byte(1)
										ptr = ptr - 2
									}
									ptr = ptr - 1
								}
								ptr = ptr + 1
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 2
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 13
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 11
								}
								ptr = ptr - 2
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 2
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 2
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr - 13
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr + 11
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 2
							}
							ptr = ptr - 12
						}
					}
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 4
				}
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
				}
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 7
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 7
					}
					ptr = ptr - 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 7
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 5
					}
					ptr = ptr + 9
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 4
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 13
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
								for buffer[ptr] != 0 {
									buffer[ptr] = buffer[ptr] - byte(1)
									ptr = ptr + 2
									buffer[ptr] = buffer[ptr] + byte(1)
									ptr = ptr - 2
								}
								ptr = ptr - 1
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 13
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
							}
							ptr = ptr - 2
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 13
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 11
							}
							ptr = ptr - 1
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 2
						}
						ptr = ptr - 12
					}
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 6
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 3
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(15)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 5
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
							}
							ptr = ptr - 2
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 2
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 4
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 9
						}
						ptr = ptr - 8
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
					}
					ptr = ptr + 9
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 9
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 9
						}
						ptr = ptr - 10
					}
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 6
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 6
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 4
					}
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
					}
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr - 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 3
							}
							ptr = ptr - 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
								ptr = ptr + 3
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr + 1
								buffer[ptr] = buffer[ptr] + byte(1)
								ptr = ptr - 4
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 9
						}
						ptr = ptr - 8
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
					}
					ptr = ptr + 9
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 9
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 9
						}
						ptr = ptr - 11
					}
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr - 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 6
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 6
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 4
					}
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 3
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 36
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 36
					}
					ptr = ptr + 5
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(15)
				for buffer[ptr] != 0 {
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 9
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 9
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 21
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 3
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 3
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 3
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 3
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 4
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 4
						}
						ptr = ptr - 4
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 4
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 13
							for buffer[ptr] != 0 {
								ptr = ptr - 9
							}
							ptr = ptr + 4
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 5
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 12
							for buffer[ptr] != 0 {
								ptr = ptr - 9
							}
							ptr = ptr + 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 6
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr - 8
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 2
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
				}
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 4
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 2
				}
				ptr = ptr + 2
			}
			ptr = ptr - 2
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 4
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 4
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 4
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 4
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 6
				fmt.Print(string(buffer[ptr]))
				ptr = ptr + 2
			}
			ptr = ptr + 4
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 7
				fmt.Print(string(buffer[ptr]))
				ptr = ptr + 7
			}
			ptr = ptr - 3
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 3
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 3
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 4
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(11)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
			}
			ptr = ptr + 4
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 9
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 14
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 7
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 7
			}
			ptr = ptr - 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 7
				buffer[ptr] = buffer[ptr] + byte(1)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr + 9
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
					}
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 7
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
						ptr = ptr + 7
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 3
					}
					ptr = ptr - 10
				}
			}
			ptr = ptr + 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 7
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 7
			}
			ptr = ptr - 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 7
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 4
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
					}
					ptr = ptr - 4
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 4
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 4
					}
					ptr = ptr + 8
				}
				ptr = ptr - 2
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 7
				for buffer[ptr] != 0 {
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
					}
					ptr = ptr - 14
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 9
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 7
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 7
					}
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 7
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 3
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 7
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 7
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 7
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 7
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 2
				for buffer[ptr] != 0 {
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 2
					}
					ptr = ptr + 4
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 7
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							ptr = ptr - 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 3
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 3
						}
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
						ptr = ptr + 7
					}
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 6
					}
					ptr = ptr - 1
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(5)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr + 9
				}
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 5
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 9
				for buffer[ptr] != 0 {
					ptr = ptr + 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 5
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 5
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 5
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 7
						}
						ptr = ptr - 7
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 7
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 16
							for buffer[ptr] != 0 {
								ptr = ptr - 9
							}
							ptr = ptr + 4
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 5
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr + 1
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 7
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 7
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 7
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 7
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 2
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr - 5
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 5
						}
						ptr = ptr - 5
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
							ptr = ptr + 5
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 14
							for buffer[ptr] != 0 {
								ptr = ptr - 9
							}
							ptr = ptr + 3
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr + 6
							for buffer[ptr] != 0 {
								ptr = ptr + 9
							}
							ptr = ptr + 1
							for buffer[ptr] != 0 {
								buffer[ptr] = buffer[ptr] - byte(1)
							}
							buffer[ptr] = buffer[ptr] + byte(1)
							ptr = ptr - 1
						}
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 1
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr - 8
					}
					ptr = ptr + 8
				}
				ptr = ptr - 9
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
				ptr = ptr + 4
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 3
				buffer[ptr] = buffer[ptr] + byte(5)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 9
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 9
					}
					ptr = ptr + 9
				}
				ptr = ptr + 4
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr - 5
				for buffer[ptr] != 0 {
					ptr = ptr - 9
				}
			}
			ptr = ptr + 3
		}
		ptr = ptr - 4
		fmt.Print(string(buffer[ptr]))
		ptr = ptr + 10
		for buffer[ptr] != 0 {
			ptr = ptr + 6
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 3
		}
		ptr = ptr - 9
		for buffer[ptr] != 0 {
			ptr = ptr - 9
		}
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(10)
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 9
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 9
		}
		ptr = ptr + 5
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 9
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 15
		for buffer[ptr] != 0 {
			ptr = ptr - 9
		}
		ptr = ptr + 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr - 8
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 8
		}
		ptr = ptr - 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 8
			buffer[ptr] = buffer[ptr] + byte(1)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + 9
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr + 8
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 7
				}
				ptr = ptr - 7
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 8
					for buffer[ptr] != 0 {
						ptr = ptr - 9
					}
					ptr = ptr + 8
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
					}
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 2
				}
				ptr = ptr - 10
			}
		}
		ptr = ptr + 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr - 8
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 8
		}
		ptr = ptr - 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 8
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 5
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
				}
				ptr = ptr - 5
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 5
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 5
				}
				ptr = ptr + 8
			}
			ptr = ptr - 1
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 8
			for buffer[ptr] != 0 {
				ptr = ptr + 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
				}
				ptr = ptr - 15
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 9
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 8
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 8
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
					}
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 7
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 7
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 8
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr - 5
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 3
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr - 8
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 8
		}
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr - 8
		for buffer[ptr] != 0 {
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 8
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 1
			for buffer[ptr] != 0 {
				ptr = ptr + 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 2
				}
				ptr = ptr + 3
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] - byte(1)
				ptr = ptr + 8
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 8
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr + 1
					for buffer[ptr] != 0 {
						ptr = ptr - 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 2
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 2
					}
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
					ptr = ptr + 8
				}
				ptr = ptr - 7
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 7
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 7
				}
				ptr = ptr - 1
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 9
			}
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(5)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
			}
			ptr = ptr + 5
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr + 27
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 6
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 9
			for buffer[ptr] != 0 {
				ptr = ptr + 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 6
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 6
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 6
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 6
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 8
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 8
					}
					ptr = ptr - 8
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 8
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 17
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
						ptr = ptr + 4
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 5
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr + 1
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 8
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 8
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 8
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr - 8
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 8
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 2
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr - 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
					}
					ptr = ptr - 6
					for buffer[ptr] != 0 {
						buffer[ptr] = buffer[ptr] - byte(1)
						ptr = ptr + 6
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 15
						for buffer[ptr] != 0 {
							ptr = ptr - 9
						}
						ptr = ptr + 3
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr + 6
						for buffer[ptr] != 0 {
							ptr = ptr + 9
						}
						ptr = ptr + 1
						for buffer[ptr] != 0 {
							buffer[ptr] = buffer[ptr] - byte(1)
						}
						buffer[ptr] = buffer[ptr] + byte(1)
						ptr = ptr - 1
					}
				}
				buffer[ptr] = buffer[ptr] + byte(1)
				ptr = ptr + 1
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr - 1
					for buffer[ptr] != 0 {
						ptr = ptr + 9
					}
					ptr = ptr - 8
				}
				ptr = ptr + 8
			}
			ptr = ptr - 9
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
			ptr = ptr + 4
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
			}
			ptr = ptr - 3
			buffer[ptr] = buffer[ptr] + byte(5)
			for buffer[ptr] != 0 {
				buffer[ptr] = buffer[ptr] - byte(1)
				for buffer[ptr] != 0 {
					buffer[ptr] = buffer[ptr] - byte(1)
					ptr = ptr + 9
					buffer[ptr] = buffer[ptr] + byte(1)
					ptr = ptr - 9
				}
				ptr = ptr + 9
			}
			ptr = ptr + 5
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr + 27
			buffer[ptr] = buffer[ptr] - byte(1)
			ptr = ptr - 6
			for buffer[ptr] != 0 {
				ptr = ptr - 9
			}
		}
		ptr = ptr + 3
	}
}
