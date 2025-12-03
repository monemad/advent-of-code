def main():
    dial_position = 50
    zero_counter = 0

    with open('day_1_input') as f:
        for rotation in f:
            direction, value = rotation[0], int(rotation[1:])
            dial_position = (dial_position + value) % 100 if direction == 'R' else (dial_position - value) % 100
            if dial_position < 0: dial_position = -dial_position % 100
            if dial_position == 0: zero_counter+=1
    
    print('password', zero_counter)

if __name__ == '__main__':
    main()