def main():
    dial_position = 50
    zero_counter = 0

    with open('day_1_input') as f:
        for rotation in f:
            direction, value = rotation[0], int(rotation[1:])
            zero_counter = zero_counter + (value // 100)
            value = value % 100
            if direction == 'R' and dial_position + value >= 100:
                zero_counter += 1
            if direction == 'L' and value >= dial_position:
                if dial_position != 0: zero_counter += 1
            dial_position = (dial_position + value) % 100 if direction == 'R' else (dial_position - value) % 100
            if dial_position < 0: dial_position = -dial_position % 100
    
    print('password', zero_counter)

if __name__ == '__main__':
    main()